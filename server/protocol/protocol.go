/*
My Protocol:

default format:
[body_size(u8)][type(u8)][body([]u8)]

handshake:
rq: [body_size(u8)][0][client_group]
rp: [body_size(u8)][0][ok]

keepalive:
rq: [body_size(u8)][1][0]
rp: [body_size(u8)][1][ok]

proxy list:
rp: [body_size(u8)][2][ip:port ip:port...]
*/
package protocol

import "fmt"

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
	output := make([]uint8, len(d.Body)+2)

	output[0] = uint8(len(d.Body)) // body size
	output[1] = uint8(d.DataType)  // data type
	copy(output[2:], d.Body)

	return output
}

// size is stripped before the data is passed
func Decode(data []uint8) (*Data, error) {
	if len(data) < 1 {
		return nil, fmt.Errorf("Data is too short: %d", len(data))
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
