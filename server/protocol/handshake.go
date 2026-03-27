package protocol

import "fmt"

func Handshake(data []uint8) (string, []byte, error) {
	d, err := Decode(data)
	if err != nil {
		return "", nil, err
	}

	if d.DataType != typeHandshake {
		return "", nil, fmt.Errorf("invalid data type: %d", d.DataType)
	}

	group := string(d.Body)
	if group == "" {
		group = "unknown"
	}

	return group, buildOK(typeHandshake), nil
}
