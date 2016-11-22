package gate

import (
	"errors"
	"time"

	"github.com/corego/tools"

	proto "github.com/aiyun/gomqtt/mqtt/protocol"
	"github.com/aiyun/gomqtt/mqtt/service"

	"fmt"

	rpc "github.com/aiyun/gomqtt/proto"
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

	// validate the user
	ok = userValidate(ci.cp.Username(), ci.cp.Password())
	if !ok {
		return errors.New("validate failed"), proto.ErrIdentifierRejected
	}

	// check mutex login
	mutexLogin(ci)

	// connect to stream
	ip, err := consist.Get(tools.Bytes2String(ci.cp.Username()))
	if err != nil {
		return err, proto.ErrServerUnavailable
	}

	ci.rpc, ok = rpcRoutes[ip]
	if !ok {
		return fmt.Errorf("no stream rpc available: %v", ip), proto.ErrServerUnavailable
	}

	err = ci.rpc.login(&rpc.LoginMsg{
		An:  ci.cp.Username(),
		Un:  ci.cp.ClientId(),
		Cid: ci.id,
		Gip: tools.String2Bytes(ci.c.LocalAddr().String()),
	})
	if err != nil {
		return err, proto.ErrServerUnavailable
	}

	// if keepalive == 0 ,we should specify a default keepalive
	if ci.cp.KeepAlive() == 0 || ci.cp.KeepAlive() < 15 {
		ci.cp.SetKeepAlive(Conf.Mqtt.DefaultKeepalive)
	}

	return nil, proto.ConnectionAccepted
}
