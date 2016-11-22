package service

import (
	"sync"
)

type conIDs struct {
	sync.RWMutex
	cids map[int64]*accountMsg
}

type accountMsg struct {
	acc   *Account
	appID []byte
}

func newconIDs() *conIDs {
	conIDs := &conIDs{
		cids: make(map[int64]*accountMsg),
	}
	return conIDs
}

func (cids *conIDs) add(cid int64, acc *accountMsg) {
	cids.Lock()
	cids.cids[cid] = acc
	cids.Unlock()
}

func (cids *conIDs) get(cid int64) (*accountMsg, bool) {
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
