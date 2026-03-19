package client

import (
	"fmt"
	"server/logger"
	"server/protocol"
)

func (c *Client) handler(group, uuid string) {
	defer remove(group, uuid)
	for {
		data, err := protocol.Read(c.Conn)
		if err != nil {
			logger.NewError(fmt.Errorf("Client read: %s", err.Error()))
			return
		}

		response, err := protocol.KeepAlive(data)
		if err != nil {
			logger.NewError(fmt.Errorf("Client keepalive: %s", err.Error()))
			return
		}

		c.Mu.Lock()
		_, err = c.Conn.Write(response)
		c.Mu.Unlock()
		if err != nil {
			logger.NewError(fmt.Errorf("Client write: %s", err.Error()))
			return
		}
	}
}
