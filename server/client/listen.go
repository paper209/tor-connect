package client

import (
	"log"
	"net"
)

func Listen() error {
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
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

		go initial(conn)
	}
}
