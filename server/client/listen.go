package client

import (
	"log"
	"net"
)

func Listen(addr string) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("[server-error] accept error: %s\n", err.Error())
			continue
		}

		go Initial(conn)
	}
}
