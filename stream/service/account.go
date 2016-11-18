package service

import "sync"

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

// GetSubUser 获取子用户
func (ats *Accounts) GetUser(acName string, uName string) (*User, bool) {
	ats.RLock()
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

// User 子用户
type User struct {
	ConV      int    // 连接版本号
	Gip       string //网关地址
	ApnsToken string //apns token
	Oline     bool   //是否在线
}

func NewUser() *User {
	user := &User{}
	return user
}

// UpdateGip 更新网关地址
// func (user *User) Update(am *proto.AccMsg, isLine bool) {
// }
