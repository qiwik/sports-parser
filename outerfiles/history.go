//Package outerfiles created for work with support files
package outerfiles

import (
	"github/sports-parser/errorpack"
	"github/sports-parser/parsingsite"
	"os"
)

//OpenHistory opens history.txt
func OpenHistory(landURL, sportsTag string, logFile *os.File) {
	historyFile, err := os.OpenFile("history.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	errorpack.ErrorErr(err, logFile)
	logFile.WriteString("History opened. Parsing is starting...\n")
	defer historyFile.Close()
	parsingsite.ParsingSports(landURL, sportsTag, historyFile, logFile)
}
