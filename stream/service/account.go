package service

import (
	"sync"
	"time"

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
	ats.RLock()
	acc, ok := ats.Accounts[string(msg.An)]
	ats.RUnlock()
	if ok {
		// Update user
		acc.Login(msg)
	} else {
		// 数据库中拉取
		// New Account
		acc := NewAccount()
		acc.Login(msg)

		ats.Lock()
		ats.Accounts[string(msg.An)] = acc
		ats.Unlock()
	}
	return nil
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
	user.ConV = msg.Cid
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
	user.ConV = msg.Cid
	user.Oline = ONLINE
	user.LastLogin = time.Now().Unix()
	acc.Unlock()
	return nil
}

// User 子用户
type User struct {
	ConV      int64  // 连接版本号
	Gip       []byte // 网关地址
	Oline     bool   // 是否在线
	LastLogin int64  // 最后登录时间
	ApnsToken []byte // apns token
}

func NewUser() *User {
	user := &User{}
	return user
}

// UpdateGip 更新网关地址
// func (user *User) Update(am *proto.AccMsg, isLine bool) {
// }
