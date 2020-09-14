package main

import (
	"fmt"
	"github/sports-parser/newstag"
	"github/sports-parser/outerfiles"
)

const (
	landURL = "https://www.sports.ru/topnews/"
)

func main() {
	inputTag := newstag.GetTag()
	fileExist := outerfiles.Revision()
	if fileExist == true {
		outerfiles.OpenHistory(landURL, inputTag)
	} else {
		outerfiles.CreateHistory(landURL, inputTag)
	}
	fmt.Scanln()
}
