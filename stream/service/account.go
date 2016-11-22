package service

import (
	"bytes"
	"sync"
	"time"

	"fmt"

	"github.com/aiyun/gomqtt/proto"
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
func (ats *Accounts) Login(msg *proto.LoginMsg) error {
	ats.Lock()
	var err error
	acc, ok := ats.Accounts[string(msg.An)]
	if ok {
		err = acc.Login(msg)
	} else {
		// 数据库中拉取
		// New Account
		acc := NewAccount()
		acc.Login(msg)
		ats.Accounts[string(msg.An)] = acc
	}
	ats.Unlock()
	return err
}

// Logout 登出
func (ats *Accounts) Logout(msg *proto.LogoutMsg) error {
	ats.Lock()
	var err error
	acc, ok := ats.Accounts[string(msg.An)]
	if ok {
		err = acc.Logout(msg)
	} else {
		err = fmt.Errorf("unfind user, an is %s, un is %s", string(msg.An), string(msg.Un))
	}
	ats.Unlock()
	return err
}

// Subscribe 订阅
func (ats *Accounts) Subscribe(msg *proto.SubMsg) error {
	ats.Lock()
	var err error
	acc, ok := ats.Accounts[string(msg.An)]
	if ok {
		err = acc.Subscribe(msg)
	} else {
		err = fmt.Errorf("unfind user, an is %s, un is %s", string(msg.An), string(msg.Un))
	}
	ats.Unlock()
	return err
}

// UnSubscribe 取消订阅
func (ats *Accounts) UnSubscribe(msg *proto.UnSubMsg) error {
	ats.Lock()
	var err error
	acc, ok := ats.Accounts[string(msg.An)]
	if ok {
		err = acc.UnSubscribe(msg)
	} else {
		err = fmt.Errorf("unfind user, an is %s, un is %s", string(msg.An), string(msg.Un))
	}
	ats.Unlock()
	return err
}

// GetSubUser 获取子用户
func (ats *Accounts) GetUser(acName string, uName string) (*User, bool) {
	// ats.RLock()
	return nil, false
}

// GetAccount 获取根用户
func (ats *Accounts) GetAccount(uname string) (*Account, bool) {

	return nil, false
}

type Account struct {
	sync.RWMutex
	Users map[string]*User //子用户
}

func NewAccount() *Account {
	account := &Account{
		Users: make(map[string]*User),
	}
	return account
}

func (acc *Account) NewUser(msg *proto.LoginMsg) error {
	user := NewUser()
	user.Gip = msg.Gip
	user.Cid = msg.Cid
	user.Oline = ONLINE
	user.LastLogin = time.Now().Unix()
	acc.Users[string(msg.Un)] = user
	return nil
}

func (acc *Account) Login(msg *proto.LoginMsg) error {
	acc.Lock()
	user, ok := acc.Users[string(msg.Un)]
	if !ok {
		user := NewUser()
		acc.Users[string(msg.Un)] = user
	}
	user.Gip = msg.Gip
	user.Cid = msg.Cid
	user.Oline = ONLINE
	user.LastLogin = time.Now().Unix()
	acc.Unlock()
	return nil
}

func (acc *Account) Logout(msg *proto.LogoutMsg) error {
	acc.Lock()
	user, ok := acc.Users[string(msg.Un)]
	if !ok {
		acc.Unlock()
		return fmt.Errorf("unfind %s, %s", string(msg.An), string(msg.Un))
	}
	if user.Cid != msg.Cid {
		acc.Unlock()
		return fmt.Errorf("user's Cid diff, old cid is %d, get cid is %d", user.Cid, msg.Cid)
	}
	user.Oline = OFFLINE
	user.LastLogout = time.Now().Unix()
	acc.Unlock()
	return nil
}

// Subscribe
func (acc *Account) Subscribe(msg *proto.SubMsg) error {
	acc.Lock()
	user, ok := acc.Users[string(msg.Un)]
	if !ok {
		acc.Unlock()
		return fmt.Errorf("unfind %s, %s", string(msg.An), string(msg.Un))
	} else {
		if user.Cid != msg.Cid {
			acc.Unlock()
			return fmt.Errorf("user's Cid diff, old cid is %d, get cid is %d", user.Cid, msg.Cid)
		}
		user.Topics = msg.Ts
	}
	acc.Unlock()
	return nil
}

// UnSubscribe
func (acc *Account) UnSubscribe(msg *proto.UnSubMsg) error {
	acc.Lock()
	user, ok := acc.Users[string(msg.Un)]
	if !ok {
		acc.Unlock()
		return fmt.Errorf("unfind %s, %s", string(msg.An), string(msg.Un))
	} else {
		if user.Cid != msg.Cid {
			acc.Unlock()
			return fmt.Errorf("user's Cid diff, old cid is %d, get cid is %d", user.Cid, msg.Cid)
		}
		// delete topics from user's topics
		for _, unSubtopic := range msg.Ts {
			for index, topic := range user.Topics {
				if bytes.Equal(unSubtopic, topic) {
					user.Topics = append(user.Topics[:index], user.Topics[index+1:]...)
				}
			}
		}
	}
	acc.Unlock()
	return nil
}

// User 子用户
type User struct {
	Cid        int64    // 连接版本号
	Gip        []byte   // 网关地址
	Oline      bool     // 是否在线
	LastLogin  int64    // 最后登录时间
	LastLogout int64    // 最后登出时间
	ApnsToken  []byte   // apns token
	Topics     [][]byte // topic列表
}

func NewUser() *User {
	user := &User{}
	return user
}
