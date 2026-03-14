package proxy

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	px "golang.org/x/net/proxy"
)

type Response struct {
	IsTor bool `json:"IsTor"`
}

const checkAPI string = "https://check.torproject.org/api/ip"

func isTor(proxy string) (bool, error) {
	dialer, err := px.SOCKS5("tcp", proxy, nil, px.Direct)
	if err != nil {
		return false, fmt.Errorf("Proxy Error: %s", err.Error())
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, proxy string) (net.Conn, error) {
				return dialer.Dial(network, proxy)
			},
		},
	}

	resp, err := client.Get(checkAPI)
	if err != nil {
		return false, fmt.Errorf("Check Proxy Error: %s", err.Error())
	}
	defer resp.Body.Close()

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return false, fmt.Errorf("Check Proxy Error: %s", err.Error())
	}

	return response.IsTor, nil
}

func checkProxies(proxies []string) []string {
	var result []string
	for _, p := range proxies {
		stat, err := isTor(p)
		if err != nil {
			log.Println(err)
			continue
		}

		if stat {
			result = append(result, p)
		}
	}

	return result
}
