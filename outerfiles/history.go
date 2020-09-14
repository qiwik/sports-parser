//Package outerfiles created for work with support files
package outerfiles

import (
	"github/sports-parser/errorpack"
	"github/sports-parser/parsingsite"
	"os"
)

//OpenHistory opens history.txt
func OpenHistory(landURL, sportsTag string) {
	historyFile, err := os.OpenFile("history.txt", os.O_APPEND|os.O_WRONLY, 0600)
	errorpack.ErrorErr(err)
	defer historyFile.Close()
	parsingsite.ParsingSports(landURL, sportsTag, historyFile)
}

//CreateHistory create new history.txt if file doesn't exist in the root
func CreateHistory(landURL, sportsTag string) {
	historyFile, err := os.Create("history.txt")
	errorpack.ErrorErr(err)
	defer historyFile.Close()
	parsingsite.ParsingSports(landURL, sportsTag, historyFile)
}
