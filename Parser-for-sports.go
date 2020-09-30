package main

import (
	"fmt"
	"github/sports-parser/logger"
	"github/sports-parser/newstag"
	"github/sports-parser/outerfiles"
)

const (
	landURL = "https://www.sports.ru/topnews/"
)

func main() {
	logFile := logger.LogInit()
	inputTag := newstag.GetTag(logFile)
	outerfiles.OpenHistory(landURL, inputTag, logFile)
	logFile.WriteString("File closed correctly.\n--------\n")
	logFile.Close()
	fmt.Scanln()
}
