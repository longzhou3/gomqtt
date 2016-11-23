package gate

import (
	"errors"
	"fmt"
	"time"

	proto "github.com/aiyun/gomqtt/mqtt/protocol"
	"github.com/aiyun/gomqtt/mqtt/service"
	"github.com/corego/tools"
)

func initConnection(ci *connInfo) (error, proto.ConnackCode) {
	// wait for connect and init connection
	// the first packet is connect type,we need to restrain the read deadline
	ci.c.SetReadDeadline(time.Now().Add(10 * time.Second))

	pt, _, _, err := service.ReadPacket(ci.c)
	if err != nil {
		code, ok := err.(proto.ConnackCode)
		if ok {
			return errors.New("read connnect packet error"), code
		}
		return err, proto.ErrIdentifierRejected
	}

	cp, ok := pt.(*proto.ConnectPacket)
	if !ok {
		return errors.New("invalid packet type"), proto.ErrIdentifierRejected
	}

	ci.cp = cp

	// transfer account and user
	accTrans(ci)

	// validate the user
	ok = userValidate(ci.acc, ci.cp.Password())
	if !ok {
		return errors.New("validate failed"), proto.ErrIdentifierRejected
	}

	// connect to stream
	ip, err := consist.Get(tools.Bytes2String(ci.acc))
	if err != nil {
		return err, proto.ErrServerUnavailable
	}

	ci.rpc, ok = rpcRoutes[ip]
	if !ok {
		return fmt.Errorf("no stream rpc available: %v", ip), proto.ErrServerUnavailable
	}

	// if keepalive == 0 ,we should specify a default keepalive
	kp := ci.cp.KeepAlive()
	switch {
	case kp == 0:
		ci.cp.SetKeepAlive(Conf.Mqtt.DefaultKeepalive)

	case kp < Conf.Mqtt.MinKeepalive:
		ci.cp.SetKeepAlive(Conf.Mqtt.MinKeepalive)

	case kp > Conf.Mqtt.MaxKeepalive:
		ci.cp.SetKeepAlive(Conf.Mqtt.MaxKeepalive)

	default:
	}

	return nil, proto.ConnectionAccepted
}
