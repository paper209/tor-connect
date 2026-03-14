package client

import (
	"fmt"
	"server/protocol"
)

func SendProxies(proxies []string) {
	payload := protocol.BuildProxyList(proxies)
	for _, group := range Clients {
		for uuid, conn := range group {
			fmt.Println(uuid) // debug
			conn.Write(payload)
		}
	}
}
