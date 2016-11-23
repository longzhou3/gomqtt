package service

import (
	"sync"

	"github.com/corego/tools"
)

type btCache struct {
	sync.RWMutex
	bts map[string]*accMsg
}

func (bc *btCache) get(topic []byte) (*accMsg, bool) {
	bc.RLock()
	if at, ok := bc.bts[string(topic)]; ok {
		bc.RUnlock()
		return at, ok
	}
	bc.RUnlock()
	return nil, false
}

func (bc *btCache) insert(acc []byte, topic []byte, appid []byte) error {
	bc.Lock()
	if acc, ok := bc.bts[string(topic)]; ok {
		acc.appID[string(appid)] = true
	} else {
		acc := newaccMsg()
		acc.appID[string(appid)] = true
		bc.bts[string(topic)] = acc
	}
	bc.Unlock()

	return nil
}

func (bc *btCache) delete(acc []byte, topic []byte, appid []byte) error {
	bc.Lock()
	if acc, ok := bc.bts[string(topic)]; ok {
		delete(acc.appID, tools.Bytes2String(appid))
	}
	bc.Unlock()
	return nil
}

func newbtCache() *btCache {
	return &btCache{bts: make(map[string]*accMsg)}
}

type accMsg struct {
	acc   []byte
	appID map[string]bool
}

func newaccMsg() *accMsg {
	return &accMsg{appID: make(map[string]bool)}
}
