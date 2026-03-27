package protocol

import "fmt"

func KeepAlive(data []uint8) ([]byte, error) {
	d, err := Decode(data)
	if err != nil {
		return nil, err
	}

	if d.DataType != typeKeepAlive {
		return nil, fmt.Errorf("invalid data type: %d", d.DataType)
	}

	return buildOK(typeKeepAlive), nil
}
