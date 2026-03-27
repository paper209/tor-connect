package client

import (
	"net"
	"sync"

	"github.com/google/uuid"
)

type Client struct {
	Conn net.Conn
	Mu   sync.Mutex
}

var (
	Clients   = make(map[string]map[string]*Client)
	ClientsMu = sync.Mutex{}
)

func new(group string, conn net.Conn) (string, *Client) {
	ClientsMu.Lock()
	defer ClientsMu.Unlock()

	groups := Clients[group]
	if groups == nil {
		groups = make(map[string]*Client)
	}

	uuid := uuid.NewString()
	groups[uuid] = &Client{
		Conn: conn,
		Mu:   sync.Mutex{},
	}

	Clients[group] = groups

	return uuid, groups[uuid]
}

func remove(group, uuid string) {
	ClientsMu.Lock()
	defer ClientsMu.Unlock()

	groups := Clients[group]
	if groups == nil {
		return
	}

	c := groups[uuid]
	if c == nil {
		return
	}
	c.Conn.Close()

	delete(groups, uuid)

	if len(groups) < 1 {
		delete(Clients, group)
	}
}
