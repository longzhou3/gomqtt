package gate

import (
	"errors"

	proto "github.com/aiyun/gomqtt/mqtt/protocol"
	"github.com/aiyun/gomqtt/mqtt/service"

	"fmt"

	rpc "github.com/aiyun/gomqtt/proto"
)

func subscribe(ci *connInfo, p *proto.SubscribePacket) error {
	topics, rets, err := topicsAndRets(ci, p.Topics(), p.Qos())
	if err != nil {
		return err
	}

	err = ci.rpc.subscribe(&rpc.SubMsg{
		Cid: ci.id,
		Ts:  topics,
	})
	if err != nil {
		return fmt.Errorf("subscribe error: %v", err)
	}

	// give back the suback
	pb := proto.NewSubackPacket()
	pb.SetPacketID(p.PacketID())

	// return the final qos level
	pb.AddReturnCodes(rets)
	service.WritePacket(ci.c, pb)

	return nil
}

func unsubscribe(ci *connInfo, p *proto.UnsubscribePacket) error {
	topics, err := topics(p)
	if err != nil {
		return err
	}

	err = ci.rpc.unSubscribe(&rpc.UnSubMsg{
		Cid: ci.id,
		Ts:  topics,
	})
	if err != nil {
		return fmt.Errorf("unSubscribe error: %v", err)
	}

	pb := proto.NewUnsubackPacket()
	pb.SetPacketID(p.PacketID())

	service.WritePacket(ci.c, pb)
	return nil
}

func topicsAndRets(ci *connInfo, tps [][]byte, qoses []byte) ([]*rpc.Topic, []byte, error) {
	rets := make([]byte, 0, len(tps))
	topics := make([]*rpc.Topic, 0, len(tps))

	for i, t := range tps {
		tp, ty, err := topicTrans(t)
		if err != nil {
			return nil, nil, err
		}

		// set master topic
		if ty == 1000 {
			// get appid and payload proto type
			err = appidTrans(ci, tp)
			if err != nil {
				return nil, nil, err
			}
		}

		var qos byte
		if qoses[i] > Conf.Mqtt.QosMax {
			qos = Conf.Mqtt.QosMax
		} else {
			qos = qoses[i]
		}

		topic := &rpc.Topic{
			Qos: int32(qos),
			Tp:  tp,
			Ty:  int32(ty),
		}

		topics = append(topics, topic)
		rets = append(rets, qos)
	}

	// 主topic必须在第一次订阅提供，因此这里需要验证主topic是否存在
	if ci.appID == nil {
		return nil, nil, errors.New("need provide master topic")
	}

	return topics, rets, nil
}

func topics(p *proto.UnsubscribePacket) ([]*rpc.Topic, error) {
	topics := make([]*rpc.Topic, 0, len(p.Topics()))

	for _, t := range p.Topics() {
		tp, ty, err := topicTrans(t)
		if err != nil {
			return nil, err
		}

		// forbid to unsubsribe master topic
		if ty == 1000 {
			return nil, errors.New("forbid to unsubsribe master topic")
		}
		topic := &rpc.Topic{
			Tp: tp,
		}

		topics = append(topics, topic)
	}

	return topics, nil
}
