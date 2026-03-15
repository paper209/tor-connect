package client

import (
	"server/protocol"
)

func SendProxies(proxies []string) {
	payload := protocol.BuildProxyList(proxies)
	for _, group := range Clients {
		for _, c := range group {
			c.Mu.Lock()
			c.Conn.Write(payload)
			c.Mu.Unlock()
		}
	}
}
