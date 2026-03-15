package client

import (
	"log"
	"server/protocol"
)

func (c *Client) handler(group, uuid string) {
	defer remove(group, uuid)
	for {
		data, err := protocol.Read(c.Conn)
		if err != nil {
			log.Printf("[client-error] read error: %s\n", err.Error())
			return
		}

		response, err := protocol.KeepAlive(data)
		if err != nil {
			log.Printf("[client-error] keepalive error: %s", err.Error())
			return
		}

		c.Mu.Lock()
		_, err = c.Conn.Write(response)
		c.Mu.Unlock()
		if err != nil {
			log.Printf("[client-error] write error: %s\n", err.Error())
			return
		}
	}
}
