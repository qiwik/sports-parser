package outerfiles

import (
	"log"
	"os"
	"time"
)

func LogInit() *os.File {
	logFile, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	currentTime := time.Now().Format("01-02-2006 15:04:05")
	logFile.WriteString("Start application: " + currentTime + "\n")
	return logFile
}
