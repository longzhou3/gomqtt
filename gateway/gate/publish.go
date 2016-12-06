package gate

/* Mqtt协议Publish报文处理模块 */
import (
	"errors"

	"github.com/aiyun/gomqtt/global"
	proto "github.com/aiyun/gomqtt/mqtt/protocol"
	"github.com/aiyun/gomqtt/uuid"
	"github.com/corego/tools"
	"github.com/uber-go/zap"

	"bytes"

	"fmt"

	rpc "github.com/aiyun/gomqtt/proto"
)

//@ToDo
//从客户端过来的只能是私聊或者单播
func publish(ci *connInfo, p *proto.PublishPacket) error {
	ci.inCount++

	switch ci.payloadProtoType {
	case global.PayloadJson:
		c2s := &global.C2SMsg{}
		err := c2s.UnmarshalJSON(p.Payload())
		if err != nil {
			return fmt.Errorf("unmarshal error: %v, data: %s", err, p.Payload())
		}

		var mid string
		if c2s.MsgID == "" {
			mid = uuid.GenStr()
		} else {
			mid = c2s.MsgID
		}

		err := ci.rpc.pubText(&rpc.PubTextMsg{
			Cid:   ci.id,
			ToAcc: tools.String2Bytes(c2s.Acc),
			Ttp:   tools.String2Bytes(c2s.Topic),
			Qos:   int32(c2s.Qos),
			Mid:   tools.String2Bytes(mid),
			Msg:   c2s.Msg,
		})
		if err != nil {
			return fmt.Errorf("pubJson rpc error: %v", err)
		}

	case global.PayloadText:
		// text格式，需要生成MsgID
		mid := tools.String2Bytes(uuid.GenStr())
		tps := bytes.Split(p.Topic(), topicSep)
		if len(tps) != 2 {
			return errors.New("invalid publish topic, text topic need to be topic--acc")
		}

		qos := qosTrans(p.QoS())

		Logger.Debug("client publish text", zap.String("topic", string(tps[0])),
			zap.String("acc", string(tps[1])), zap.Int("in_count", ci.inCount))
		err := ci.rpc.pubText(&rpc.PubTextMsg{
			Cid:   ci.id,
			ToAcc: tps[1],
			Ttp:   tps[0],
			Qos:   int32(qos),
			Mid:   mid,
			Msg:   p.Payload(),
		})
		if err != nil {
			return fmt.Errorf("pubText rpc error: %v", err)
		}

	case global.PayloadProtobuf:

	}

	// need give back the ack
	if p.QoS() >= 1 {
		pb := proto.NewPubackPacket()
		pb.SetPacketID(p.PacketID())
		write(ci, pb)
	}

	return nil
}

// 将消息ID反映射后，投递到stream删除
func puback(ci *connInfo, p *proto.PubackPacket) error {
	mid := p.PacketID()

	ci.rwm.RLock()
	ids, ok := ci.idMap[mid]
	ci.rwm.RUnlock()
	if ok {
		err := ci.rpc.puback(&rpc.PubAckMsg{
			Acc:  ci.acc,
			Plty: ci.payloadProtoType,
			Mids: ids,
		})
		if err != nil {
			return fmt.Errorf("puback rpc error : %v", err)
		}

		ci.rwm.Lock()
		delete(ci.idMap, mid)
		ci.rwm.Unlock()
	}

	return nil
}
