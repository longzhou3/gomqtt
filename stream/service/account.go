package service

import (
	"fmt"
	"sync"
	"time"

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

// Login 登陆
func (ats *Accounts) Login(msg *proto.LoginMsg) (*Account, error) {
	ats.Lock()
	var err error
	acc, ok := ats.Accounts[string(msg.An)]
	if ok {
		err = acc.Login(msg)
	} else {
		// 数据库中拉取
		// New Account
		acc = NewAccount()
		acc.Login(msg)
		ats.Accounts[string(msg.An)] = acc
	}
	ats.Unlock()
	return acc, err
}

// Logout 登出
func (ats *Accounts) Logout(msg *proto.LogoutMsg) error {

	// ats.Lock()
	// var err error
	// acc, ok := ats.Accounts[string(msg.An)]
	// if ok {
	// 	err = acc.Logout(msg)
	// } else {
	// 	err = fmt.Errorf("unfind user, an is %s, un is %s", string(msg.An), string(msg.Un))
	// }
	// ats.Unlock()
	// return err
	return nil
}

// Subscribe 订阅
func (ats *Accounts) Subscribe(msg *proto.SubMsg) error {
	// ats.Lock()
	// var err error
	// acc, ok := ats.Accounts[string(msg.An)]
	// if ok {
	// 	err = acc.Subscribe(msg)
	// } else {
	// 	err = fmt.Errorf("unfind user, an is %s, un is %s", string(msg.An), string(msg.Un))
	// }
	// ats.Unlock()
	// return err
	return nil
}

// UnSubscribe 取消订阅
func (ats *Accounts) UnSubscribe(msg *proto.UnSubMsg) error {
	// ats.Lock()
	// var err error
	// acc, ok := ats.Accounts[string(msg.An)]
	// if ok {
	// 	err = acc.UnSubscribe(msg)
	// } else {
	// 	err = fmt.Errorf("unfind user, an is %s, un is %s", string(msg.An), string(msg.Un))
	// }
	// ats.Unlock()
	// return err
	return nil
}

// GetSubUser 获取子用户
func (ats *Accounts) GetUser(acName string, uName string) (*AppID, bool) {
	// ats.RLock()
	return nil, false
}

// GetAccount 获取根用户
func (ats *Accounts) GetAccount(uname string) (*Account, bool) {

	return nil, false
}

type Account struct {
	sync.RWMutex
	AppIDs map[string]*AppID //子用户
}

func NewAccount() *Account {
	account := &Account{
		AppIDs: make(map[string]*AppID),
	}
	return account
}

func (acc *Account) NewUser(msg *proto.LoginMsg) error {
	appID := NewAppID()
	appID.Gip = msg.Gip
	appID.Cid = msg.Cid
	appID.Oline = ONLINE
	appID.LastLogin = time.Now().Unix()
	acc.AppIDs[tools.Bytes2String(msg.AppID)] = appID
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
	appID.Oline = ONLINE
	appID.LastLogin = time.Now().Unix()

	// 订阅
	for _, topic := range msg.Ts {
		appID.Topics[string(topic.Tp)] = topic
	}

	acc.Unlock()
	return nil
}

func (acc *Account) Logout(un []byte) error {
	acc.Lock()
	appID, ok := acc.AppIDs[tools.Bytes2String(un)]
	if !ok {
		acc.Unlock()
		return fmt.Errorf("unfind appID %s", tools.Bytes2String(un))
	}
	appID.Oline = OFFLINE
	appID.LastLogout = time.Now().Unix()
	acc.Unlock()
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
	// if msg.AppID != nil {
	// 	delete(acc.AppIDs, tools.Bytes2String(un))
	// 	acc.AppIDs[string(msg.AppID)] = appID
	// }
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

// AppID appid
type AppID struct {
	Cid        int64  // 连接版本号
	Gip        []byte // 网关地址
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
