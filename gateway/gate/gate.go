package gate

import (
	"fmt"
	"time"
)

type Gate struct {
}

func New() *Gate {
	return &Gate{}
}

func (g *Gate) Start(isStatic bool) {
	// init configurations
	fmt.Println(isStatic)

	loadConfig(isStatic)

	time.Sleep(6 * time.Second)

	// init providers
	providersStart()

	// init addmin service
	go adminStart()

	// start the monitors
	// monitorsStart()

	Logger.Info("gate started!!")
}
