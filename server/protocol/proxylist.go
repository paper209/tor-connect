package protocol

func BuildProxyList() []uint8 {
	data := &Data{
		DataType: typeProxyList,
		Body:     "127.0.0.1:8080 127.0.0.1:2323",
	}

	return data.Encode()
}
