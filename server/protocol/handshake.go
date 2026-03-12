package protocol

import "fmt"

// return values: group, response, error
func Handshake(data []uint8) (string, []byte, error) {
	d, err := Decode(data)
	if err != nil {
		return "", nil, err
	}

	// check data type
	if d.DataType != typeHandshake {
		return "", nil, fmt.Errorf("invalid data type: %d", d.DataType)
	}

	group := string(d.Body)
	if group == "" {
		group = "unknown"
	}

	return group, buildOK(typeHandshake), nil
}
