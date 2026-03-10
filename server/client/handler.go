package client

import (
	"log"
	"net"
	"server/protocol"
)

func handler(group, uuid string, conn net.Conn) {
	defer remove(group, uuid)
	for {
		data, err := protocol.Read(conn)
		if err != nil {
			log.Printf("[client-error] read error: %s\n", err.Error())
			return
		}

		response, err := protocol.KeepAlive(data)
		if err != nil {
			log.Printf("[client-error] keepalive error: %s", err.Error())
			return
		}

		_, err = conn.Write(response)
		if err != nil {
			log.Printf("[client-error] write error: %s\n", err.Error())
			return
		}
	}
}
