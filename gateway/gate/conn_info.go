package gate

import (
	"net"
	"sync"

	proto "github.com/aiyun/gomqtt/mqtt/protocol"
)

type connInfo struct {
	id int64
	c  net.Conn
	cp *proto.ConnectPacket

	inCount  int
	outCount int

	stopped chan struct{}

	relogin chan struct{}

	rpc *rpcServie

	test []byte
}

type connInfos struct {
	sync.RWMutex
	infos map[int64]*connInfo
}

var cons = &connInfos{
	infos: make(map[int64]*connInfo),
}

func saveCI(ci *connInfo) {
	cons.Lock()
	cons.infos[ci.id] = ci
	cons.Unlock()
}

func getCI(id int64) *connInfo {
	cons.RLock()
	c, ok := cons.infos[id]
	cons.RUnlock()

	if ok {
		return c
	}

	return nil
}

func delCI(id int64) {
	cons.Lock()
	delete(cons.infos, id)
	cons.Unlock()
}
