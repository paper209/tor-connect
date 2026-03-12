package client

import (
	"fmt"
	"server/protocol"
)

func SendProxyList() {
	for _, group := range Clients {
		for uuid, conn := range group {
			fmt.Println(uuid)
			conn.Write(protocol.BuildProxyList())
		}
	}
}
