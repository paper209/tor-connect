package main

import (
	"log"
	"server/admin"
	"server/client"
	"server/config"
	"server/proxy"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Println(err.Error())
		return
	}

	go proxy.Handler(cfg)
	go client.Listen(cfg.Client.ListenAddr)

	admin.Handler()
}
