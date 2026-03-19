package proxy

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"server/logger"
	"sync"
	"time"

	px "golang.org/x/net/proxy"
)

type Response struct {
	IsTor bool `json:"IsTor"`
}

const checkAPI string = "https://check.torproject.org/api/ip"

func isTor(addr string) (bool, error) {
	dialer, err := px.SOCKS5("tcp", addr, nil, px.Direct)
	if err != nil {
		return false, fmt.Errorf("Proxy connect: %s", err.Error())
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
		return false, fmt.Errorf("Proxy request: %s", err.Error())
	}
	defer resp.Body.Close()

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return false, fmt.Errorf("Proxy decode: %s", err.Error())
	}

	return response.IsTor, nil
}

func checkProxy(wg *sync.WaitGroup, mu *sync.Mutex, result *[]string, addr string) {
	defer wg.Done()

	stat, err := isTor(addr)
	if err != nil {
		log.Println(err)
		return
	}

	if stat {
		mu.Lock()
		*result = append(*result, addr)
		mu.Unlock()
	}
}

func checkProxies(proxies []string) []string {
	logger.NewINFO("Proxy check started")

	var (
		result []string
		mu     sync.Mutex
		wg     sync.WaitGroup
	)

	for _, p := range proxies {
		wg.Add(1)
		go checkProxy(&wg, &mu, &result, p)
	}

	wg.Wait()
	return result
}
