package gate

import (
	"strconv"

	"github.com/corego/tools"

	"fmt"

	proto "github.com/aiyun/gomqtt/mqtt/protocol"
	"github.com/aiyun/gomqtt/mqtt/service"
	rpc "github.com/aiyun/gomqtt/proto"
)

// login and first subscribe
func login(ci *connInfo, p *proto.SubscribePacket) error {
	// set app id and topics
	topics, rets, err := topicsAndRets(ci, p)
	if err != nil {
		return err
	}

	// check mutex login
	mutexLogin(ci)

	// 这里的AppID先设置为CLientId，具体参见connInfo结构
	err = ci.rpc.login(&rpc.LoginMsg{
		An:    ci.acc,
		AppID: ci.appID,
		Cid:   ci.id,
		Gip:   tools.String2Bytes(ci.c.LocalAddr().String()),
		Ts:    topics,
	})
	if err != nil {
		return err
	}

	// subscribe the cid topic in nats
	cstr := strconv.FormatInt(ci.id, 10)
	_, err = nc.Subscribe(cstr, sub2nats)
	return fmt.Errorf("sub to nats error: %v", err)

	// give back the suback
	pb := proto.NewSubackPacket()
	pb.SetPacketID(p.PacketID())

	// return the final qos level
	pb.AddReturnCodes(rets)
	service.WritePacket(ci.c, pb)
	return nil
}
