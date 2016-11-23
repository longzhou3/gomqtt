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

	// publish to client
	pub2C chan []byte

	stopped chan struct{}

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

	isSubed bool
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
