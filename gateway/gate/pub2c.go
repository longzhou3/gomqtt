package gate

import (
	"bytes"
	"fmt"

	"github.com/aiyun/gomqtt/global"
	"github.com/nats-io/nats"
	"github.com/uber-go/zap"

	proto "github.com/aiyun/gomqtt/mqtt/protocol"
	"github.com/aiyun/gomqtt/mqtt/service"
)

func pub2c(ci *connInfo, msg *nats.Msg) {
	switch ci.payloadProtoType {
	case global.PayloadProtobuf:

	case global.PayloadText:
		d := &global.TextMsgs{}
		_, err := d.UnmarshalMsg(msg.Data)
		if err != nil {
			Logger.Info("unmarshal TextMsgs error", zap.Error(err), zap.Int64("cid", ci.id))
			return
		}

		err = pubText(ci, d)
		if err != nil {
			Logger.Debug("pubText error", zap.Error(err))
		}

	case global.PayloadJson:
	}

}

func pubText(ci *connInfo, m *global.TextMsgs) error {
	for _, msg := range m.Msgs {
		p := proto.NewPublishPacket()
		p.SetQoS(byte(msg.Qos))
		//@Optimize
		topic := bytes.Join([][]byte{msg.FTopic, msg.FAcc}, topicSep)
		p.SetTopic(topic)
		p.SetPayload(msg.Msg)

		id, err := mapID(ci, [][]byte{msg.MsgID}, msg.Qos)
		if err != nil {
			return err
		}

		p.SetPacketID(id)

		Logger.Debug("recv nats msg", zap.String("msg", string(msg.Msg)), zap.Int("id", int(id)))
		err = service.WritePacket(ci.c, p)
		if err != nil {
			return err
		}
	}
	return nil
}

func mapID(ci *connInfo, ids [][]byte, qos int32) (uint16, error) {
	id, err := getID(ci)
	if err != nil {
		return 0, err
	}

	// 只有在qos不为0时，才保存id映射，为后面ack的删除做备用
	if qos != 0 {
		ci.idMap[id] = ids
	}

	return id, nil
}

// 自加MsgID,超过65535返回错误
func getID(ci *connInfo) (uint16, error) {
	if ci.msgID >= 65535 {
		return 0, fmt.Errorf("MsgID Beyond the maximum range")
	}
	ci.msgID++
	return ci.msgID, nil
}
