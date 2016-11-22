package gate

import (
	"sync"

	"github.com/corego/tools"
)

type account struct {
	ci    *connInfo
	users map[string]*connInfo
}

var accounts = make(map[string]*account)
var al = &sync.RWMutex{}

func mutexLogin(ci *connInfo) {
	switch Conf.Mutex.Type {
	case 1: // mutex login disable

	case 2: // mutex with account(Username)
		al.RLock()
		acc, ok := accounts[string(ci.cp.Username())]
		al.RUnlock()

		if ok {
			// kick off the former connections
			acc.ci.stopped <- struct{}{}

			// wait for kick off done
			<-acc.ci.relogin
		}

		//update now one
		al.Lock()
		accounts[string(ci.cp.Username())] = &account{
			ci: ci,
		}
		al.Unlock()

	case 3: //mutex with account + user(Username + clientID)
		al.RLock()
		acc, ok := accounts[string(ci.cp.Username())]
		al.RUnlock()

		if ok {
			c, ok := acc.users[string(ci.cp.ClientId())]
			if ok {
				c.stopped <- struct{}{}

				<-c.relogin
			}

			acc.users[string(ci.cp.ClientId())] = ci
		} else {
			accounts[string(ci.cp.Username())] = &account{
				users: map[string]*connInfo{
					tools.Bytes2String(ci.cp.ClientId()): ci,
				},
			}
		}

	case 4: //mutex with clientid
		al.RLock()
		acc, ok := accounts[string(ci.cp.ClientId())]
		al.RUnlock()

		if ok {
			// kick off the former connections
			acc.ci.stopped <- struct{}{}

			// wait for kick off done
			<-acc.ci.relogin
		}

		//update now one
		al.Lock()
		accounts[string(ci.cp.ClientId())] = &account{
			ci: ci,
		}
		al.Unlock()

		return
	}

}

func delMutex(ci *connInfo) {
	switch Conf.Mutex.Type {
	case 1:

	case 2:
		al.Lock()
		delete(accounts, tools.Bytes2String(ci.cp.Username()))
		al.Unlock()

	case 3:
		al.Lock()
		acc, _ := accounts[string(ci.cp.Username())]
		delete(acc.users, tools.Bytes2String(ci.cp.ClientId()))
		al.Unlock()
	case 4:
		al.Lock()
		delete(accounts, tools.Bytes2String(ci.cp.ClientId()))
		al.Unlock()
	}
}
