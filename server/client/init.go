package client

import (
	"fmt"
	"net"
	"server/logger"
	"server/protocol"
)

func Initial(conn net.Conn) {
	data, err := protocol.Read(conn)
	if err != nil {
		conn.Close()
		logger.NewError(fmt.Errorf("client read error: %s", err.Error()))
		return
	}

	group, response, err := protocol.Handshake(data)
	if err != nil {
		conn.Close()
		logger.NewError(fmt.Errorf("client handshake error: %s", err.Error()))
		return
	}

	_, err = conn.Write(response)
	if err != nil {
		conn.Close()
		logger.NewError(fmt.Errorf("client write error: %s", err.Error()))
		return
	}

	uuid, c := new(group, conn)
	go c.handler(group, uuid)
}
