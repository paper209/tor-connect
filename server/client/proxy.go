package client

import (
	"fmt"
	"server/logger"
	"server/protocol"
)

func SendProxies(proxies []string) {
	ClientsMu.Lock()
	defer ClientsMu.Unlock()

	count := 0
	payload := protocol.BuildProxyList(proxies)
	for _, group := range Clients {
		for _, c := range group {
			c.Mu.Lock()
			c.Conn.Write(payload)
			c.Mu.Unlock()

			count++
		}
	}

	logger.NewINFO(fmt.Sprintf("proxy distribution completed: %d", count))
}
