package main

import (
	"server/admin"
	"server/client"
	"server/config"
	"server/logger"
	"server/proxy"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		logger.NewError(err)
		return
	}

	go proxy.Handler(cfg)
	go client.Listen(cfg.Client.ListenAddr)

	admin.Handler()
}
