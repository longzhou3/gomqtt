package gate

import (
	"time"

	"github.com/aiyun/gomqtt/global"
	"github.com/nats-io/nats"
	"github.com/uber-go/zap"
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

func sub2nats(msg *nats.Msg) {
	m := &global.Pub2C{}
	_, err := m.UnmarshalMsg(msg.Data)
	if err != nil {
		Logger.Info("unmarshal nats data error", zap.Error(err))
		return
	}

	ci := getCI(m.Cid)
	if ci == nil {
		Logger.Info("cant find specify cid", zap.Int64("cid", m.Cid))
		return
	}

	select {
	case ci.pub2C <- m:

	default:
		Logger.Info("send to pub2C failed", zap.Int("pub2C_len", len(ci.pub2C)), zap.Int64("cid", m.Cid))
	}
}

func pub2c(msg *global.Pub2C) {

}
