package service

import (
	"time"

	"github.com/nats-io/nats"
	"github.com/uber-go/zap"
)

type natsInfo struct {
	nc *nats.Conn // connect nats cluster
}

// newnatsInfo return *natsInfo
func newnatsInfo(addrs []string) (*natsInfo, error) {
	// init natsInfo
	nc, err := initnatsInfo(addrs)
	if err != nil {
		return nil, err
	}
	natsInfo := natsInfo{nc: nc}
	return &natsInfo, nil
}

func initnatsInfo(addrs []string) (*nats.Conn, error) {
	nc := initNatsConn(addrs)
	return nc, nil
}

// push 推送消息至nats服务
// func (ns *natsInfo) push(subject string, pm *basic.PushMessage) error {
// 	msgpStart := time.Now()
// 	b, err := pm.MarshalMsg(nil)
// 	if err != nil {
// 		return err
// 	}

// 	natsStart := time.Now()
// 	err = ns.nc.Publish(subject, b)

// 	return err
// }

// apnspush apnspush
// func (ns *natsInfo) apnspush(subject string, data []byte) error {
// 	natsStart := time.Now()
// 	err := ns.nc.Publish(subject, data)
// 	natsTu := time.Now().Sub(natsStart).Nanoseconds()
// 	natsTimer.Update(time.Duration(natsTu))
// 	natsGauge.Update(natsTu)
// 	return err
// }

// apnspush apnspush
// func (ns *natsInfo) statisticsPush(subject string, data []byte) error {
// 	natsStart := time.Now()
// 	err := ns.nc.Publish(subject, data)
// 	natsTu := time.Now().Sub(natsStart).Nanoseconds()
// 	natsTimer.Update(time.Duration(natsTu))
// 	natsGauge.Update(natsTu)
// 	return err
// }

func initNatsConn(addrs []string) *nats.Conn {
	opts := nats.DefaultOptions
	opts.Servers = addrs
	opts.MaxReconnect = 100
	opts.ReconnectWait = 20 * time.Second
	// opts.NoRandomize = true

	nc, err := opts.Connect()
	if err != nil {
		Logger.Panic("Nats", zap.String("Connect", err.Error()))
		return nil
	}

	// Setup callbacks to be notified on disconnects and reconnects
	nc.Opts.DisconnectedCB = func(nc *nats.Conn) {
		Logger.Info("Nats", zap.String("ConnectedUrl", nc.ConnectedUrl()))
	}

	// See who we are connected to on reconnect.
	nc.Opts.ReconnectedCB = func(nc *nats.Conn) {
		Logger.Info("Nats", zap.String("ConnectedUrl", nc.ConnectedUrl()))
	}

	return nc
}
