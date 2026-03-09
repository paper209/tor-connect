package protocol

// my protocol: [body_size(u8)][type(u8)][body([]u8)]
type Data struct{}

func (d *Data) Encode() []uint8 {
	return nil
}

func Decode(data []uint8) (*Data, error) {
	return nil, nil
}
