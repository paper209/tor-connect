package protocol

import (
	"encoding/binary"
	"fmt"
)

type dataType uint8

const (
	typeHandshake dataType = 0
	typeKeepAlive dataType = 1
	typeProxyList dataType = 2
)

type Data struct {
	DataType dataType
	Body     string
}

func (d *Data) Encode() []uint8 {
	output := make([]uint8, len(d.Body)+3)

	binary.BigEndian.PutUint16(output[0:2], uint16(len(d.Body)))
	output[2] = uint8(d.DataType)
	copy(output[3:], d.Body)

	return output
}

func Decode(data []uint8) (*Data, error) {
	if len(data) < 1 {
		return nil, fmt.Errorf("data is too short: %d", len(data))
	}

	return &Data{
		DataType: dataType(data[0]),
		Body:     string(data[1:]),
	}, nil
}

func buildOK(t dataType) []uint8 {
	data := &Data{
		DataType: t,
		Body:     "ok",
	}

	return data.Encode()
}
