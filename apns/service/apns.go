package service

import "github.com/nats-io/nats"

type Apns struct {
}

func New() *Apns {
	return &Apns{}
}

// nats connection
var nc *nats.Conn

func (g *Apns) Start(isStatic bool) {
	loadConfig(isStatic)

	// tempory disable reload
	// go adminStart()
}
