package gate

import (
	"fmt"
	"time"

	"github.com/aiyun/gomqtt/global"
	"github.com/nats-io/nats"
	"github.com/uber-go/zap"

	"bytes"

	proto "github.com/aiyun/gomqtt/mqtt/protocol"
	"github.com/aiyun/gomqtt/mqtt/service"
)

func initNatsConn() (*nats.Conn, error) {
	opts := nats.DefaultOptions
	opts.Servers = Conf.Gateway.NatsAddrs
	opts.MaxReconnect = 1000
	opts.ReconnectWait = 5 * time.Second

	nc, err := opts.Connect()
	if err != nil {
		return nil, err
	}

	// Setup callbacks to be notified on disconnects and reconnects
	nc.Opts.DisconnectedCB = func(nc *nats.Conn) {
		Logger.Error("nats disconnected")
	}

	// See who we are connected to on reconnect.
	nc.Opts.ReconnectedCB = func(nc *nats.Conn) {
		Logger.Info("nats reconnected")
	}

	return nc, nil
}

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

	case global.PayloadJson:
	}

}

func pubText(ci *connInfo, m *global.TextMsgs) error {
	for _, msg := range m.Msgs {
		p := &proto.PublishPacket{}
		p.SetQoS(byte(msg.Qos))
		//@Optimize
		topic := bytes.Join([][]byte{msg.FTopic, msg.FAcc}, topicSep)
		p.SetTopic(topic)
		p.SetPayload(msg.Msg)

		err := mapID(ci, [][]byte{msg.MsgID})
		if err != nil {
			return err
		}

		service.WritePacket(ci.c, p)
	}
	return nil
}

func mapID(ci *connInfo, ids [][]byte) error {
	id, err := getID(ci)
	if err != nil {
		return err
	}

	ci.idMap[id] = ids
	return nil
}

// 自加MsgID,超过65535返回错误
func getID(ci *connInfo) (uint16, error) {
	if ci.msgID >= 65535 {
		return 0, fmt.Errorf("MsgID Beyond the maximum range")
	}
	ci.msgID++
	return ci.msgID, nil
}
