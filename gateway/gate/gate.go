package gate

import (
	"time"

	"github.com/aiyun/gomqtt/uuid"
)

type Gate struct {
}

func New() *Gate {
	return &Gate{}
}

func (g *Gate) Start(isStatic bool) {
	// init configurations
	loadConfig(isStatic)

	// wait for rpc connections inited
	time.Sleep(6 * time.Second)

	// start uuid service
	uuid.Init(Conf.Gateway.ServerId, "ts16", Logger)

	// init providers
	providersStart()

	// init addmin service
	go adminStart()

	// start the monitors
	// monitorsStart()

	Logger.Info("gate started!!")
}
