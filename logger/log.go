package logger

import (
	"log"
	"os"
	"time"
)

//LogInit creates or opens a log file in the root
func LogInit() *os.File {
	logFile, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal("Error: can't work with log.txt")
	}
	currentTime := time.Now().Format("01-02-2006 15:04:05")
	logFile.WriteString("Start application: " + currentTime + "\n")
	return logFile
}

//LogErrorClose closes a log file while an error exists
func LogErrorClose(logFile *os.File, err string) {
	currentTime := time.Now().Format("01-02-2006 15:04:05")
	logFile.WriteString(err + " " + currentTime + "\n------\n")
	logFile.Close()
}
