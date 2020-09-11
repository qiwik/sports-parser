package main

import (
	"github/sports-parser/newstag"
	"github/sports-parser/parsingsite"
)

const (
	landURL = "https://www.sports.ru/topnews/"
)

func main() {
	inputTag := newstag.GetTag()
	parsingsite.ParsingSports(landURL, inputTag)
}
