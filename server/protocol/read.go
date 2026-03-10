package protocol

import (
	"io"
	"net"
)

func Read(conn net.Conn) ([]byte, error) {
	size := make([]byte, 1)
	_, err := io.ReadFull(conn, size)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, 1+size[0])
	_, err = io.ReadFull(conn, buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
