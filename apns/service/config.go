package service

import (
	"io/ioutil"
	"log"
	"time"

	"fmt"

	"github.com/naoina/toml"
	"github.com/nats-io/nats"
	"github.com/uber-go/zap"
)

type Config struct {
	Common struct {
		Version  string
		IsDebug  bool
		LogLevel string
		LogPath  string
	}

	Apns struct {
		NatsAddrs []string
	}
}

var Conf = &Config{}

func loadConfig(staticConf bool) {
	var contents []byte
	var err error

	if staticConf {
		//静态配置
		contents, err = ioutil.ReadFile("configs/apns.toml")
	} else {
		contents, err = ioutil.ReadFile("/etc/gomqtt/apns.toml")
	}

	if err != nil {
		log.Fatal("load config error", zap.Error(err))
	}

	tbl, err := toml.Parse(contents)
	if err != nil {
		log.Fatal("parse config error", zap.Error(err))
	}

	toml.UnmarshalTable(tbl, Conf)

	fmt.Println(Conf)

	// 初始化Logger
	InitLogger(Conf.Common.LogPath, Conf.Common.LogLevel, Conf.Common.IsDebug)

	nc, err = initNatsConn()
	if err != nil {
		Logger.Fatal("init nats error", zap.Error(err))
	}
}

func initNatsConn() (*nats.Conn, error) {
	opts := nats.DefaultOptions
	opts.Servers = Conf.Apns.NatsAddrs
	opts.MaxReconnect = 1000
	opts.ReconnectWait = 5 * time.Second

	nc, err := opts.Connect()
	if err != nil {
		return nil, err
	}

	// Setup callbacks to be notified on disconnects and reconnects
	nc.Opts.DisconnectedCB = func(nc *nats.Conn) {
		Logger.Error("nats disconnected")
	}

	// See who we are connected to on reconnect.
	nc.Opts.ReconnectedCB = func(nc *nats.Conn) {
		Logger.Info("nats reconnected")
	}

	return nc, nil
}
