package logger

import (
	"log"
	"os"
)

//LogInit creates or opens a log file in the root
func LogInit() (*log.Logger, *os.File) {
	logFile, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal("Error: can't work with log.txt")
	}
	loggerFile := log.New(logFile, "sports-parser ", log.LstdFlags|log.Lshortfile)
	loggerFile.Println("Start application")
	return loggerFile, logFile
}
