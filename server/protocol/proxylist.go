package protocol

func BuildProxyList(proxies []string) []uint8 {
	var body string
	for _, proxy := range proxies {
		if len(body)+len(proxy)+1 > 248 {
			break
		}
		body += proxy + " "
	}

	data := &Data{
		DataType: typeProxyList,
		Body:     body,
	}

	return data.Encode()
}
