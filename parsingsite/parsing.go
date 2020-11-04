//Package parsingsite allows you to get data from the site
package parsingsite

import (
	"fmt"
	"net/http"
	"os"
	"sports-parser/errorpack"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

//ParsingSports parses html and write data to the standart output
func ParsingSports(landURL string, sportsTag []string, historyFile, logFile *os.File) {
	var err interface{}

	accessFlag := false //Verifies that there is information related to tag
	var numberOfNews int

	if len(sportsTag) == 1 {
		historyFile.WriteString("Current tag: " + sportsTag[0] + "\n\n")
	} else {
		historyFile.WriteString("Current tags: ")
		for i := range sportsTag {
			historyFile.WriteString(sportsTag[i])
		}
		historyFile.WriteString("\n\n")
	}

	getSite, err := http.Get(landURL) //Request to the site
	if err != nil {
		num := 1
		errorpack.ErrorHandler(logFile, err, num)
	}

	topnews, err := goquery.NewDocumentFromReader(getSite.Body) //Read html of the main page
	if err != nil {
		num := 2
		errorpack.ErrorHandler(logFile, err, num)
	}

	topnews.Find(".short-news").Each(func(_ int, shortNews *goquery.Selection) { //Parsing of the big list of the news
		var err interface{}

		date := shortNews.Find("b").Text()           //Date of the news
		times, err := shortNews.Find(".time").Html() //News release time
		if err != nil {
			num := 2
			errorpack.ErrorHandler(logFile, err, num)
		}

		shortNews.Find(".short-text").Each(func(_ int, shortText *goquery.Selection) { //Parsing of the short news step by step
			var err interface{}

			newsURL, err := shortText.Attr("href") //URL of each news
			if err != true {
				num := 2
				errorpack.ErrorHandler(logFile, err, num)
			}
			if string(newsURL[0]) != "/" { //Incorrect link check
				return
			}

			fullURL := "https://www.sports.ru" + newsURL
			getNews, err := http.Get(fullURL) //Request to the news page
			if err != nil {
				num := 2
				errorpack.ErrorHandler(logFile, err, num)
			}

			news, err := goquery.NewDocumentFromReader(getNews.Body) //Read html of the separate news page
			if err != nil {
				num := 2
				errorpack.ErrorHandler(logFile, err, num)
			}

			title := news.Find(".h1_size_tiny").Text() //News title

			news.Find(".news-item__tags-line").Each(func(_ int, tagItems *goquery.Selection) { //Parsing of the tags list
				arrayOfTags := []string{} //Tags from site

				tagItems.Find(".link_size_small").Each(func(_ int, tagItem *goquery.Selection) { //Parsing of the tags step by step
					var err interface{}

					tag, err := tagItem.Attr("title") //Tag from site
					arrayOfTags = append(arrayOfTags, tag)
					if err != true {
						num := 2
						errorpack.ErrorHandler(logFile, err, num)
					}
				})

				coincidence := len(sportsTag) //Number of matches required for news
				for i := range sportsTag {
					for j := range arrayOfTags {
						if strings.ToLower(sportsTag[i]) == strings.ToLower(arrayOfTags[j]) {
							coincidence--
							break
						}
					}
				}

				if coincidence == 0 { //If all incoming tags are in the news so
					accessFlag = true

					currentDate := time.Now().Format("01-02-2006 15:04:05")

					historyFile.WriteString("Time of the request: " + currentDate + "\n")
					historyFile.WriteString("News:" + title + "\n")
					historyFile.WriteString(date + times + " Link:" + fullURL + "\n")
					historyFile.WriteString("-----" + "\n")
					numberOfNews++

					logFile.WriteString("Successful reading.\n")

					fmt.Printf("Time of the request: %s\nNews:%s\n%s%s Link: %s\n-----\n", currentDate, title, date, times, fullURL)
				}
			})
		})
	})
	historyFile.WriteString("\n***\n")
	fmt.Printf("Total news: %d\n", numberOfNews)
	fmt.Println("Search finished. Press enter to quit.")
	logFile.WriteString("Search finished " + time.Now().Format("01-02-2006 15:04:05") + "\n")
	if accessFlag == false {
		fmt.Println("The tag hasn't news related to it. Please, restart the application.") //Information related to tag doesn't exist
		logFile.WriteString("Incorrect tag.\n")
	}
}
