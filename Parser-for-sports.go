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
	loggerFile, file := logger.LogInit()
	inputTag := newstag.GetTag(loggerFile, file)
	outerfiles.OpenHistory(landURL, inputTag, loggerFile, file)
	loggerFile.Println("File closed correctly.")
	file.Close()
	fmt.Scanln()
}
