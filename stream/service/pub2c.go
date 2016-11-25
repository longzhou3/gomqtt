package service

import (
	"log"
	"runtime/debug"
	"strconv"
	"time"

	"github.com/aiyun/gomqtt/global"
	"github.com/aiyun/gomqtt/proto"
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

// pushText 推送消息至nats服务
func (ns *natsInfo) pushText(subject string, msg *global.TextMsgs) error {

	b, err := msg.MarshalMsg(nil)
	if err != nil {
		return err
	}

	log.Println("pushText --- nats", subject, len(msg.Msgs))

	err = ns.nc.Publish(subject, b)
	return err
}

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

// taskMsg msg
type taskMsg struct {
	cid int64
	acc []byte
	ts  []*proto.Topic
}

func addTask(t *taskMsg) {

	log.Println("addTask", t)
	gStream.taskChan <- t
}

func startDealTask(taskn int) {

	stopChan := make(chan bool, 1)
	for index := 0; index < taskn; index++ {
		go func() {
			dealTask(gStream.taskChan, stopChan)
		}()
	}
}

func dealTask(ch chan *taskMsg, sc chan bool) {
	//异常接收代码
	defer func() {
		close(ch)
		close(sc)

		log.Println(zap.Stack())
		if err := recover(); err != nil {
			log.Printf("panic: %s\nStack trace:\n%s", err, debug.Stack())
			Logger.Error("startDealTask", zap.Stack())
		}
	}()
	for {
		select {
		case t, ok := <-ch:
			if !ok {
				Logger.Panic("Chan", zap.String("dealTask", "recv chan failed"))
				break
			}
			// Logger.Info("dealTask", zap.Object("taskMsg", t))
			log.Println("dealTask", t)
			PushOffLineMsg(t)
			break
		case <-sc:
			Logger.Info("Chan", zap.String("dealTask", "get stop signal"))
			return

		}
	}
}

func PushOffLineMsg(t *taskMsg) {
	for _, topicMsg := range t.ts {
		// topicMsg.

		msgids := gStream.cache.msgIDManger.GetMsgIDs(t.acc, topicMsg.Tp)

		if msgids != nil {
			for mmm, msgidMsg := range msgids.MsgID {
				log.Println(mmm, "----------msgidMsg : ", msgidMsg)
			}
			// push
			MsgsCacha := make([]*global.TextMsg, 0, len(msgids.MsgID))
			// get Msg
			for _, msgidMsg := range msgids.MsgID {
				if data, ok := gStream.cache.msgCache.Get(msgidMsg.MsgID); ok {
					var Qos int32
					if msgidMsg.MsgQos <= topicMsg.Qos {
						Qos = msgidMsg.MsgQos
					} else {
						Qos = topicMsg.Qos
					}
					Msg := &global.TextMsg{
						FAcc:       t.acc,
						FTopic:     topicMsg.Tp,
						RetryCount: 3,
						Qos:        Qos,
						MsgID:      msgidMsg.MsgID,
						Msg:        data,
					}
					log.Println("data is  ", string(data))
					MsgsCacha = append(MsgsCacha, Msg)
				}
			}
			pushMsg := &global.TextMsgs{
				Msgs: MsgsCacha,
			}
			gStream.nats.pushText(strconv.FormatInt(t.cid, 10), pushMsg)
		}
	}
}

// if tycid, ok := acc.STopics[string(msg.Ttp)]; ok {
// 	if tycid.nastTopic != "0" {
// 		var Qos int32
// 		if msg.Qos <= tycid.qos {
// 			Qos = msg.Qos
// 		} else {
// 			Qos = tycid.qos
// 		}
// 		Msg := &global.TextMsg{
// 			FAcc:       facc.Acc,
// 			FTopic:     msg.Ttp,
// 			RetryCount: 3,
// 			Qos:        Qos,
// 			MsgID:      msg.Mid,
// 			Msg:        msg.Msg,
// 		}
// 		msg := &global.TextMsgs{
// 			Msgs: []*global.TextMsg{Msg},
// 		}
// 		// push to nats
// 		gStream.nats.pushText(tycid.nastTopic, msg)
// 	}
// }
