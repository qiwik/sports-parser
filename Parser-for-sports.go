package main

import (
	"fmt"
	"sports-parser/logger"
	"sports-parser/newstag"
	"sports-parser/outerfiles"
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
