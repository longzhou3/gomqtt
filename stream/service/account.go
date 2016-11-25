package service

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/aiyun/gomqtt/global"
	proto "github.com/aiyun/gomqtt/proto"
	"github.com/corego/tools"
	"github.com/uber-go/zap"
)

const (
	ONLINE  bool = true  //在线
	OFFLINE bool = false //离线
)

type Accounts struct {
	sync.RWMutex
	Accounts map[string]*Account
}

func NewAccounts() *Accounts {
	as := &Accounts{
		Accounts: make(map[string]*Account),
	}
	return as
}

func (ats *Accounts) GetAccAndAppID(accid, appid []byte) (*Account, *AppID, bool) {
	ats.RLock()
	if acc, ok := ats.Accounts[string(accid)]; ok {
		if appid, ok := acc.AppIDs[string(appid)]; ok {
			ats.RUnlock()
			return acc, appid, true
		}
	}
	ats.RUnlock()
	return nil, nil, false
}

// Login 登陆
func (ats *Accounts) Login(msg *proto.LoginMsg) error {
	ats.Lock()
	var err error
	acc, ok := ats.Accounts[string(msg.Acc)]
	if ok {
		err = acc.Login(msg)
	} else {
		// 数据库中拉取
		// New Account
		acc = NewAccount(msg.Acc)
		acc.Login(msg)
		ats.Accounts[string(msg.Acc)] = acc
	}
	ats.Unlock()
	return err
}

// Logout logout
func (ats *Accounts) Logout(accid, appid []byte) error {
	ats.Lock()
	var err error
	acc, ok := ats.Accounts[string(accid)]
	if ok {
		acc.Logout(appid)
	} else {
		err = fmt.Errorf("unfind acc %s, appid %s", tools.Bytes2String(accid), tools.Bytes2String(appid))
	}
	ats.Unlock()
	return err
}

func (ats *Accounts) PubText(facc *Account, appid *AppID, msg *proto.PubTextMsg) error {
	var err error
	ats.RLock()
	acc, ok := ats.Accounts[string(msg.ToAcc)]
	ats.RUnlock()
	if ok {
		err = acc.PubText(facc, appid, msg)
	}
	return err
}

type Account struct {
	sync.RWMutex
	Acc    []byte
	AppIDs map[string]*AppID //子用户
	// @ToDo
	// 这里还要保存topic的订阅等级
	STopics map[string]*topicTypeCid
	PTopics map[string]*topicTypeCid
}

// topicTypeCid topic 的类型和订阅该topic的cid （string）类型
type topicTypeCid struct {
	topicTy   int32
	qos       int32
	nastTopic string
}

func newtopicTypeCid(ty int32, qos int32, nastTopic string) *topicTypeCid {
	return &topicTypeCid{topicTy: ty, qos: qos, nastTopic: nastTopic}
}

func NewAccount(Acc []byte) *Account {
	account := &Account{
		Acc:     Acc,
		AppIDs:  make(map[string]*AppID),
		STopics: make(map[string]*topicTypeCid),
		PTopics: make(map[string]*topicTypeCid),
	}
	return account
}

func (acc *Account) NewUser(msg *proto.LoginMsg) error {
	appID := NewAppID()
	appID.Gip = msg.Gip
	appID.Cid = msg.Cid
	appID.Oline = ONLINE
	appID.LastLogin = time.Now().Unix()
	acc.AppIDs[string(msg.AppID)] = appID
	return nil
}

func (acc *Account) Login(msg *proto.LoginMsg) error {
	acc.Lock()
	var appID *AppID
	appID, ok := acc.AppIDs[string(msg.AppID)]
	if !ok {
		appID = NewAppID()
		acc.AppIDs[string(msg.AppID)] = appID
	}
	appID.Gip = msg.Gip
	appID.Cid = msg.Cid
	appID.PT = msg.PT
	appID.Oline = ONLINE
	appID.LastLogin = time.Now().Unix()

	for _, topic := range msg.Ts {
		if global.PChatTopic == topic.Ty || global.SPushTopic == topic.Ty {
			// 保存topic 和 cid
			topic2Nats := newtopicTypeCid(topic.Ty, topic.Qos, strconv.FormatInt(msg.Cid, 10))
			acc.STopics[string(topic.Tp)] = topic2Nats
		} else if global.BPushTopic == topic.Ty {
			// 广播特殊处理
		}
	}

	for k, v := range acc.STopics {
		log.Println("查看订阅", k, v)
	}
	// 订阅
	for _, topic := range msg.Ts {
		appID.Topics[string(topic.Tp)] = topic
	}
	acc.Unlock()
	return nil
}

func (acc *Account) Logout(appid []byte) error {
	acc.Lock()
	appID, ok := acc.AppIDs[string(appid)]
	if !ok {
		acc.Unlock()
		return fmt.Errorf("unfind appID %s", tools.Bytes2String(appid))
	}
	appID.Oline = OFFLINE
	appID.LastLogout = time.Now().Unix()
	acc.Unlock()

	// cid 置为0
	for _, topic := range appID.Topics {
		if topic2Nats, ok := acc.STopics[string(topic.Tp)]; ok {
			topic2Nats.nastTopic = "0"
		}
	}
	return nil
}

// Subscribe
func (acc *Account) Subscribe(un []byte, msg *proto.SubMsg) error {
	acc.Lock()
	appID, ok := acc.AppIDs[string(un)]
	if !ok {
		acc.Unlock()
		return fmt.Errorf("unfind appID %s  ", tools.Bytes2String(un))
	} else {
		for _, topic := range msg.Ts {
			appID.Topics[string(topic.Tp)] = topic
		}
	}
	acc.Unlock()

	Logger.Info("Subscribe", zap.String("Gip", fmt.Sprintf("%s", appID.Gip)))
	for _, topic := range appID.Topics {
		Logger.Info("Subscribe", zap.String("Topic", fmt.Sprintf("%s", topic.Tp)))
	}

	return nil
}

// UnSubscribe
func (acc *Account) UnSubscribe(un []byte, msg *proto.UnSubMsg) error {
	acc.Lock()
	appID, ok := acc.AppIDs[string(un)]
	if !ok {
		acc.Unlock()
		return fmt.Errorf("unfind appID %s  ", string(un))
	} else {
		for _, topic := range msg.Ts {
			if _, ok := appID.Topics[string(topic.Tp)]; ok {
				delete(appID.Topics, tools.Bytes2String(topic.Tp))
			}
		}
	}
	acc.Unlock()

	Logger.Info("UnSubscribe", zap.String("Gip", fmt.Sprintf("%s", appID.Gip)))
	for _, topic := range appID.Topics {
		Logger.Info("UnSubscribe", zap.String("Topic", fmt.Sprintf("%s", topic.Tp)))
	}
	return nil
}

func (acc *Account) PubText(facc *Account, appid *AppID, msg *proto.PubTextMsg) error {
	if tycid, ok := acc.STopics[string(msg.Ttp)]; ok {
		if tycid.nastTopic != "0" {
			var Qos int32
			if msg.Qos <= tycid.qos {
				Qos = msg.Qos
			} else {
				Qos = tycid.qos
			}
			Msg := &global.TextMsg{
				FAcc:       facc.Acc,
				FTopic:     msg.Ttp,
				RetryCount: 3,
				Qos:        Qos,
				MsgID:      msg.Mid,
				Msg:        msg.Msg,
			}
			msg := &global.TextMsgs{
				Msgs: []*global.TextMsg{Msg},
			}
			// push to nats
			gStream.nats.pushText(tycid.nastTopic, msg)
		}
	}
	return nil
}

// AppID appid
type AppID struct {
	Cid        int64  // 连接版本号
	Gip        []byte // 网关地址
	PT         int32  // payload 协议
	Oline      bool   // 是否在线
	LastLogin  int64  // 最后登录时间
	LastLogout int64  // 最后登出时间
	ApnsToken  []byte // apns token
	Topics     map[string]*proto.Topic
	// Topics     [][]byte // topic列表
}

func NewAppID() *AppID {
	appID := &AppID{
		Topics: make(map[string]*proto.Topic),
	}
	return appID
}
