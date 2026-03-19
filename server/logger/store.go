package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"server/config"
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

func store(t logType, now, message string) {
	logs := make(map[string]Log)
	data, err := os.ReadFile(config.LogPATH)
	if err != nil {
		NewError(fmt.Errorf("Log store: %s", err.Error()))
		return
	}

	json.Unmarshal(data, &logs)
	logs[now] = Log{
		Type:    t,
		Message: message,
	}

	data, err = json.Marshal(logs)
	if err != nil {
		NewError(fmt.Errorf("Log marshal: %s", err.Error()))
		return
	}

	err = os.WriteFile(config.LogPATH, data, 0644)
	if err != nil {
		NewError(fmt.Errorf("Log write: %s", err.Error()))
		return
	}
}
