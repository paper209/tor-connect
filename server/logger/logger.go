package logger

import (
	"log"
	"os"
	"time"
)

var (
	infoLogger  = log.New(os.Stdout, "\033[32m[INFO]\033[0m ", log.LstdFlags)
	errorLogger = log.New(os.Stdout, "\033[31m[ERROR]\033[0m ", log.LstdFlags)
)

func NewINFO(message string) {
	now := time.Now().Format("2006/01/02 15:04:05")
	infoLogger.Println(message)

	store(INFO, now, message)
}

func NewError(err error) {
	now := time.Now().Format("2006/01/02 15:04:05")
	errorLogger.Println(err.Error())

	store(ERROR, now, err.Error())
}
