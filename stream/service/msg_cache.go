package service

import (
	"sync"
)

type MsgCache struct {
	sync.RWMutex
	Msgs map[string][]byte
}

func NewMsgCache() *MsgCache {
	msgcache := &MsgCache{
		Msgs: make(map[string][]byte),
	}
	return msgcache
}

func (msgCache *MsgCache) Insert(msgid string, msg []byte) error {
	msgCache.Lock()
	msgCache.Msgs[msgid] = msg
	msgCache.Unlock()
	return nil
}

func (msgCache *MsgCache) Delete(msgid string) error {
	msgCache.Lock()
	delete(msgCache.Msgs, msgid)
	msgCache.Unlock()
	return nil
}

func (msgCache *MsgCache) Get(msgid string) ([]byte, bool) {
	msgCache.RLock()
	if msg, ok := msgCache.Msgs[msgid]; ok {
		msgCache.RUnlock()
		return msg, true
	}
	msgCache.RUnlock()
	return nil, false
}

type MsgIdCache struct {
	sync.RWMutex
	AccMsg map[string]*AccMsg
}

type AccMsg struct {
	sync.RWMutex
	AppMsg map[string]*AppMsgIDs
}

func NewAccMsg() *AccMsg {
	return &AccMsg{AppMsg: make(map[string]*AppMsgIDs)}
}

type AppMsgIDs struct {

	// 单播
	// 私聊
	// 广播
}

func NewMsgIdCache() *MsgIdCache {
	return &MsgIdCache{AccMsg: make(map[string]*AccMsg)}
}
