//Package outerfiles created for work with support files
package outerfiles

import (
	"github/sports-parser/logger"
	"github/sports-parser/parsingsite"
	"log"
	"os"
)

//OpenHistory opens history.txt
func OpenHistory(landURL string, sportsTag []string, logFile *os.File) {
	historyFile, err := os.OpenFile("history.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		err := "Error: history.txt can not be created or opened"
		logger.LogErrorClose(logFile, err)
		log.Fatal(err)
	}
	logFile.WriteString("History opened. Parsing is starting...\n")
	defer historyFile.Close()
	parsingsite.ParsingSports(landURL, sportsTag, historyFile, logFile)
}
