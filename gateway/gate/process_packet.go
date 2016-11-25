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
		// 如果是通过appid管理的topic订阅方案，那么不允许再主动订阅
		if ci.isInstantLogin {
			err = errors.New("you cant subscribe any topics after instant login")
		} else {
			// 还未订阅，进行登录和订阅
			if !ci.isSubed {
				err = loginAndSub(ci, p.Topics(), p.Qos(), p.PacketID())
			} else {
				err = subscribe(ci, p)
			}
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
