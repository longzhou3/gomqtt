package gate

import (
	"strconv"

	"github.com/corego/tools"
	"github.com/nats-io/nats"

	"fmt"

	proto "github.com/aiyun/gomqtt/mqtt/protocol"
	"github.com/aiyun/gomqtt/mqtt/service"
	rpc "github.com/aiyun/gomqtt/proto"
)

// login and first subscribe
func loginAndSub(ci *connInfo, tps [][]byte, qoses []byte, pid uint16) error {
	// set app id and topics
	topics, rets, err := topicsAndRets(ci, tps, qoses)
	if err != nil {
		return err
	}

	// check mutex login
	mutexLogin(ci)

	// 这里的AppID先设置为CLientId，具体参见connInfo结构
	err = ci.rpc.login(&rpc.LoginMsg{
		Acc:   ci.acc,
		AppID: ci.appID,
		PT:    ci.payloadProtoType,
		Cid:   ci.id,
		Gip:   tools.String2Bytes(ci.c.LocalAddr().String()),
		Ts:    topics,
	})
	if err != nil {
		return fmt.Errorf("login rpc error: %v", err)
	}

	// subscribe the cid topic in nats
	cstr := strconv.FormatInt(ci.id, 10)
	h, err := nc.Subscribe(cstr, func(m *nats.Msg) {
		pub2c(ci, m)
	})
	if err != nil {
		return fmt.Errorf("sub to nats error: %v", err)
	}

	ci.natsHandler = h

	// give back the suback
	pb := proto.NewSubackPacket()
	pb.SetPacketID(pid)

	// return the final qos level
	pb.AddReturnCodes(rets)
	service.WritePacket(ci.c, pb)

	ci.isSubed = true
	return nil
}
