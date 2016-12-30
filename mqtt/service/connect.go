package service

import (
	"net"

	"github.com/taitan-io/gomqtt/mqtt/protocol"
)

func connectPacket(conn net.Conn) (*protocol.ConnectPacket, error) {
	buf, err := Read(conn)
	if err != nil {
		return nil, err
	}

	cp := protocol.NewConnectPacket()

	_, err = cp.Decode(buf)
	return cp, err
}
