package logger

import (
	"log"
	"os"
)

var (
	infoLogger  = log.New(os.Stdout, "\033[32m[INFO]\033[0m ", log.LstdFlags)
	errorLogger = log.New(os.Stdout, "\033[31m[ERROR]\033[0m ", log.LstdFlags)
)

func NewINFO(message string) {
	infoLogger.Println(message)
}

func NewError(err error) {
	errorLogger.Println(err.Error())
}
