package service

import (
	"sync"
)

// NewMsgCache   key:msgid, value:msg
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

// MsgIdCache
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
	SIDs *SPushID
	// 私聊
	PIDs *PPushID
	// 广播
	BIDs *BPushID
	// 群聊
	GIDs *GPushID
}

func NewAppMsgIDs() *AppMsgIDs {
	apms := &AppMsgIDs{}
	return apms
}

type MsgID struct {
	Qos int
}

type SPushID struct {
	IDs map[string]*MsgID
}

func NewSPushID() *SPushID {
	return &SPushID{IDs: make(map[string]*MsgID)}
}

type PPushID struct {
	IDs map[string]*MsgID
}

func NewPPushID() *PPushID {
	return &PPushID{IDs: make(map[string]*MsgID)}
}

type BPushID struct {
	IDs   []*MsgID
	Index map[string]int
}

func NewBPushID() *BPushID {
	return &BPushID{IDs: make([]*MsgID, 0), Index: make(map[string]int)}
}

type GPushID struct {
	IDs   []*MsgID
	Index map[string]int
}

func NewGPushID() *GPushID {
	return &GPushID{IDs: make([]*MsgID, 0), Index: make(map[string]int)}
}

func NewMsgIdCache() *MsgIdCache {
	return &MsgIdCache{AccMsg: make(map[string]*AccMsg)}
}
