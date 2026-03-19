package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Proxy struct {
	PATH        string `json:"proxies_path"`
	SenderDelay int    `json:"sender_delay"`
	Checker     bool   `json:"checker"`
}

type Client struct {
	ListenAddr string `json:"listen_addr"`
	Proxy      Proxy  `json:"proxy"`
}

type Config struct {
	LogPATH string `json:"log_path"`
	Client  Client `json:"client"`
}

var LogPATH string

func Read() (*Config, error) {
	data, err := os.ReadFile("config.json")
	if err != nil {
		return nil, fmt.Errorf("Config: %s", err.Error())
	}

	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, fmt.Errorf("Config: %s", err.Error())
	}

	LogPATH = cfg.LogPATH

	return &cfg, nil
}
