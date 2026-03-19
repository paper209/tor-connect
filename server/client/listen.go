package client

import (
	"fmt"
	"net"
	"server/logger"
)

func Listen(addr string) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer ln.Close()

	logger.NewINFO(fmt.Sprintf("Server is running on %s", addr))

	for {
		conn, err := ln.Accept()
		if err != nil {
			logger.NewError(fmt.Errorf("Server accept: %s", err.Error()))
			continue
		}

		go Initial(conn)
	}
}
