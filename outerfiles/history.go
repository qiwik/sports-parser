//Package outerfiles created for work with support files
package outerfiles

import (
	"log"
	"os"
	"sports-parser/errorpack"
	"sports-parser/parsingsite"
)

//OpenHistory opens history.txt
func OpenHistory(landURL string, sportsTag []string, logFile *log.Logger, file *os.File) {
	var err interface{}
	historyFile, err := os.OpenFile("history.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		num := 3
		errorpack.ErrorHandler(logFile, err, num, file)
	}
	logFile.Println("History opened. Parsing is starting...")
	defer historyFile.Close()
	parsingsite.ParsingSports(landURL, sportsTag, historyFile, logFile, file)
}
