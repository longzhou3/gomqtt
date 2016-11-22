package gate

import (
	proto "github.com/aiyun/gomqtt/mqtt/protocol"
	"github.com/aiyun/gomqtt/mqtt/service"

	"fmt"

	rpc "github.com/aiyun/gomqtt/proto"
)

func subscribe(ci *connInfo, p *proto.SubscribePacket) error {
	topics, rets := topicsAndRets(p)

	err := ci.rpc.subscribe(&rpc.SubMsg{
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
	topics := topics(p)
	err := ci.rpc.unSubscribe(&rpc.UnSubMsg{
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

func topicsAndRets(p *proto.SubscribePacket) ([]*rpc.Topic, []byte) {
	rets := make([]byte, 0, len(p.Topics()))
	topics := make([]*rpc.Topic, 0, len(p.Topics()))

	for i, t := range p.Topics() {
		var qos byte
		if p.Qos()[i] > Conf.Mqtt.QosMax {
			qos = Conf.Mqtt.QosMax
		} else {
			qos = p.Qos()[i]
		}

		topic := &rpc.Topic{
			Qos: int32(qos),
			Tp:  t,
		}

		topics = append(topics, topic)
		rets = append(rets, qos)
	}

	return topics, rets
}

func topics(p *proto.UnsubscribePacket) []*rpc.Topic {
	topics := make([]*rpc.Topic, 0, len(p.Topics()))

	for _, t := range p.Topics() {
		topic := &rpc.Topic{
			Tp: t,
		}

		topics = append(topics, topic)
	}

	return topics
}
