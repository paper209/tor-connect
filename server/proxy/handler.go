package proxy

import (
	"log"
	"server/client"
	"server/config"
	"time"
)

func Handler(cfg *config.Config) {
	delay := time.Duration(cfg.Client.Proxy.SenderDelay) * time.Second
	for {
		proxies, err := readProxies(cfg.Client.Proxy.PATH)
		if err != nil {
			log.Println(err.Error())
			time.Sleep(60 * time.Second)
			continue
		}

		proxies = checkProxies(proxies)
		client.SendProxies(proxies)

		time.Sleep(delay)
	}
}
