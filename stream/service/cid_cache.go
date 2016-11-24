package service

import (
	"sync"

	"github.com/aiyun/gomqtt/proto"
)

type conIDs struct {
	sync.RWMutex
	cids map[int64]*accMsg
}

func newconIDs() *conIDs {
	conIDs := &conIDs{
		cids: make(map[int64]*accMsg),
	}
	return conIDs
}

// accMsg 存放acccound和appid的映射关系,主题下面保存该用户的信息
type accMsg struct {
	acc   []byte
	appID []byte
}

func newaccMsg() *accMsg {
	accmsg := &accMsg{}
	return accmsg
}

func (cids *conIDs) add(msg *proto.LoginMsg) {
	cids.Lock()
	if acc, ok := cids.cids[msg.Cid]; ok {
		acc.appID = msg.AppID
	} else {
		acc := newaccMsg()
		acc.acc = msg.Acc
		acc.appID = msg.AppID
		cids.cids[msg.Cid] = acc
	}
	cids.Unlock()
}

func (cids *conIDs) get(cid int64) (*accMsg, bool) {
	cids.RLock()
	if acc, ok := cids.cids[cid]; ok {
		cids.RUnlock()
		return acc, true
	}
	cids.RUnlock()
	return nil, false
}

func (cids *conIDs) delete(cid int64) bool {
	cids.Lock()
	if _, ok := cids.cids[cid]; ok {
		delete(cids.cids, cid)
		cids.Unlock()
		return true
	}
	cids.Unlock()
	return false
}

type conID struct {
	acc *Account
}
