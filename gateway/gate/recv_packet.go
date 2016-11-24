package gate

import (
	"fmt"
	"net"
	"time"

	"github.com/aiyun/gomqtt/mqtt/service"
	"github.com/uber-go/zap"
)

//@ToDo
// 将recvPacket合并到主goroutine
func recvPacket(ci *connInfo) {
	defer func() {
		close(ci.stopped)
	}()

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

			break
		}

		err = processPacket(ci, pt)
		if err != nil {
			Logger.Info("process packet error", zap.Error(err), zap.Int64("cid", ci.id))
			break
		}

		ci.inCount++
	}
}
