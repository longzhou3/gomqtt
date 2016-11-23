package gate

import (
	"errors"

	"fmt"

	proto "github.com/aiyun/gomqtt/mqtt/protocol"
	"github.com/aiyun/gomqtt/mqtt/service"
)

func processPacket(ci *connInfo, pt proto.Packet) error {
	var err error

	switch p := pt.(type) {
	case *proto.DisconnectPacket: // recv Disconnect
		err = errors.New("recv disconnect packet")

	case *proto.PublishPacket: // recv publish
		err = publish(ci, p)

	case *proto.PubackPacket:
		err = puback(ci, p)

	case *proto.SubscribePacket:
		if !ci.isSubed {
			err = login(ci, p)
			ci.isSubed = true
		} else {
			err = subscribe(ci, p)
		}

	case *proto.UnsubscribePacket:
		err = unsubscribe(ci, p)

	case *proto.PingreqPacket:
		pingReq(ci)

	default:
		err = fmt.Errorf("recv invalid packet type: %T", pt)
	}

	return err
}

func pingReq(ci *connInfo) {
	pb := proto.NewPingrespPacket()
	service.WritePacket(ci.c, pb)
}
