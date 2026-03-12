package client

import (
	"log"
	"net"
	"server/protocol"
)

func initial(conn net.Conn) {
	data, err := protocol.Read(conn)
	if err != nil {
		conn.Close()
		log.Printf("[client-error] read error: %s", err.Error())

		return
	}

	group, response, err := protocol.Handshake(data)
	if err != nil {
		conn.Close()
		log.Printf("[client-error] handshake error: %s", err.Error())

		return
	}

	_, err = conn.Write(response)
	if err != nil {
		conn.Close()
		log.Printf("[client-error] write error: %s\n", err.Error())

		return
	}

	uuid := new(group, conn)
	go handler(group, uuid, conn)
}
