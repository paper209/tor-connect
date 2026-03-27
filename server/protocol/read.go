package protocol

import (
	"encoding/binary"
	"io"
	"net"
)

func Read(conn net.Conn) ([]byte, error) {
	size := make([]byte, 2)
	_, err := io.ReadFull(conn, size)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, 1+binary.BigEndian.Uint16(size[0:]))
	_, err = io.ReadFull(conn, buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
