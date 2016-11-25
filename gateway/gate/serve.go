package gate

import (
	"fmt"
	"net"
	"time"

	proto "github.com/aiyun/gomqtt/mqtt/protocol"

	"github.com/corego/tools"
	"github.com/uber-go/zap"

	rpc "github.com/aiyun/gomqtt/proto"
	"github.com/aiyun/gomqtt/uuid"
	"github.com/nats-io/nats"

	"github.com/aiyun/gomqtt/mqtt/service"
)

type connInfo struct {
	id int64

	c net.Conn

	cp *proto.ConnectPacket

	inCount  int
	outCount int

	relogin chan struct{}

	rpc *rpcServie

	test []byte

	acc []byte

	// appID只能是以下几种形式：
	// 1.username--appid传递的，这里的appid是在服务器做了Topics管理的(动态Topics管理)
	// 2.通过主topic type == 1000 来传递的，这里是静态类型的appid.在这种情况下，connect时首先将
	// appid设置为ClientID，然后后续subscribe时，再替换为主topic，若没有主topic，那么
	// 当前连接时异常的，必须断开
	appID []byte

	isSubed        bool
	isInstantLogin bool

	payloadProtoType int32

	msgID uint16
	idMap map[uint16][][]byte

	natsHandler *nats.Subscription
}

func serve(c net.Conn) {
	defer func() {
		if err := recover(); err != nil {
			Logger.Info("user's main goroutine has a panic error", zap.Error(err.(error)), zap.Stack())
		}
	}()

	// init a new connInfo
	ci := &connInfo{}

	//generate a uuid for this conn
	ci.id = uuid.Gen()
	ci.c = c
	Logger.Debug("a new connection has established", zap.Int64("cid", ci.id), zap.String("ip", c.RemoteAddr().String()))

	defer func() {
		c.Close()

		if ci.isSubed {
			delMutex(ci)
			err := ci.natsHandler.Unsubscribe()
			if err != nil {
				Logger.Info("unsubscribe error", zap.Error(err), zap.Int64("cid", ci.id))
			}
		}

		close(ci.relogin)
	}()

	ci.relogin = make(chan struct{})
	ci.idMap = make(map[uint16][][]byte)

	//----------------Connection init---------------------------------------------
	err := connect(ci)
	if err != nil {
		return
	}

	Logger.Debug("user connected ok!", zap.String("acc", tools.Bytes2String(ci.acc)),
		zap.String("user", tools.Bytes2String(ci.appID)), zap.String("password", tools.Bytes2String(ci.cp.Password())), zap.Int64("cid", ci.id), zap.Float64("keepalive", float64(ci.cp.KeepAlive())))

	wait := time.Duration(ci.cp.KeepAlive()+10) * time.Second
	for {
		if !ci.isSubed {
			// if not subscribed，only wait for 10 second
			ci.c.SetReadDeadline(time.Now().Add(time.Duration(Conf.Mqtt.MinKeepalive-5) * time.Second))
		} else {
			ci.c.SetReadDeadline(time.Now().Add(wait))
		}

		// We need to considering about the network delay,so here allows 10 seconds delay.
		pt, buf, n, err := service.ReadPacket(ci.c)
		if err != nil {
			nerr, ok := err.(net.Error)
			if ok && nerr.Timeout() {
				Logger.Debug("user connect but not subscribed, disconnected", zap.Int64("cid", ci.id))
			} else {
				Logger.Warn("Read packet error", zap.Error(err), zap.String("buf", fmt.Sprintf("%v", buf)), zap.Int("bytes", n), zap.Int64("cid", ci.id))
			}

			goto STOP
		}

		err = processPacket(ci, pt)
		if err != nil {
			Logger.Info("process packet error", zap.Error(err), zap.Int64("cid", ci.id))
			goto STOP
		}

		ci.inCount++
	}

	// go recvPacket(ci)

	// loop reading data
	// for {
	// 	select {
	// 	case <-ci.stopped:
	// 		Logger.Info("user's main thread is going to stop", zap.Int64("cid", ci.id))
	// 		goto STOP
	// 	}
	// }

STOP:
	if ci.isSubed {
		err = ci.rpc.logout(&rpc.LogoutMsg{
			Cid: ci.id,
		})

		Logger.Debug("user logout", zap.Error(err), zap.Int64("cid", ci.id))
	}
}
