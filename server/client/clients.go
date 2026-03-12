package client

import (
	"net"
	"sync"

	"github.com/google/uuid"
)

var (
	Clients   = make(map[string]map[string]net.Conn) // map[group]map[uuid]conn
	ClientsMu = sync.Mutex{}
)

func new(group string, conn net.Conn) string {
	ClientsMu.Lock()
	defer ClientsMu.Unlock()

	groups := Clients[group]
	if groups == nil {
		groups = make(map[string]net.Conn)
	}

	uuid := uuid.NewString()
	groups[uuid] = conn

	Clients[group] = groups

	return uuid
}

func remove(group, uuid string) {
	ClientsMu.Lock()
	defer ClientsMu.Unlock()

	groups := Clients[group]
	if groups == nil {
		return
	}

	conn := groups[uuid]
	if conn == nil {
		return
	}
	conn.Close()

	delete(groups, uuid)

	if len(groups) < 1 {
		delete(Clients, group)
	}
}
