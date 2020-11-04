//Package outerfiles created for work with support files
package outerfiles

import (
	"os"
	"sports-parser/errorpack"
	"sports-parser/parsingsite"
)

//OpenHistory opens history.txt
func OpenHistory(landURL string, sportsTag []string, logFile *os.File) {
	var err interface{}
	historyFile, err := os.OpenFile("history.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		num := 3
		errorpack.ErrorHandler(logFile, err, num)
	}
	logFile.WriteString("History opened. Parsing is starting...\n")
	defer historyFile.Close()
	parsingsite.ParsingSports(landURL, sportsTag, historyFile, logFile)
}
