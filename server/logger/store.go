package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"server/config"
	"sync"
)

type logType string

const (
	INFO  logType = "info"
	ERROR logType = "error"
)

type Log struct {
	Type    logType `json:"log_type"`
	Message string  `json:"message"`
}

var mu sync.Mutex

func store(t logType, now, message string) {
	mu.Lock()
	defer mu.Unlock()

	logs := make(map[string]Log)
	data, err := os.ReadFile(config.LogPATH)
	if err != nil {
		NewError(fmt.Errorf("log stored error: %s", err.Error()))
		return
	}

	json.Unmarshal(data, &logs)
	logs[now] = Log{
		Type:    t,
		Message: message,
	}

	data, err = json.Marshal(logs)
	if err != nil {
		NewError(fmt.Errorf("log marshal error: %s", err.Error()))
		return
	}

	err = os.WriteFile(config.LogPATH, data, 0644)
	if err != nil {
		NewError(fmt.Errorf("log write: %s", err.Error()))
		return
	}
}
