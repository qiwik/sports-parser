//Package parsingsite allows you to get data from the site
package parsingsite

import (
	"fmt"
	"github/sports-parser/errorpack"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

//ParsingSports parses html and write data to the standart output
func ParsingSports(landURL, sportsTag string, historyFile *os.File) {
	accessFlag := false //Verifies that there is information related to tag

	historyFile.WriteString("Current tag: " + sportsTag + "\n\n")
	currentDate := time.Now().Format("01-02-2006 15:04:05")

	getSite, err := http.Get(landURL) //Request to the site
	errorpack.ErrorErr(err)

	topnews, err := goquery.NewDocumentFromReader(getSite.Body) //Read html of the main page
	errorpack.ErrorErr(err)

	topnews.Find(".short-news").Each(func(_ int, shortNews *goquery.Selection) { //Parsing of the big list of the news

		date := shortNews.Find("b").Text()          //Date of the news
		time, err := shortNews.Find(".time").Html() //News release time
		errorpack.ErrorErr(err)

		shortNews.Find(".short-text").Each(func(_ int, shortText *goquery.Selection) { //Parsing of the short news step by step

			newsURL, errBool := shortText.Attr("href") //URL of each news
			errorpack.ErrorBool(errBool)
			if string(newsURL[0]) != "/" { //Incorrect link check
				return
			}

			fullURL := "https://www.sports.ru" + newsURL
			getNews, err := http.Get(fullURL) //Request to the news page
			errorpack.ErrorErr(err)

			news, err := goquery.NewDocumentFromReader(getNews.Body) //Read html of the separate news page
			errorpack.ErrorErr(err)

			title := news.Find(".h1_size_tiny").Text() //News title

			news.Find(".news-item__tags-line").Each(func(_ int, tagItems *goquery.Selection) { //Parsing of the tags list
				tagItems.Find(".link_size_small").Each(func(_ int, tagItem *goquery.Selection) { //Parsing of the tags step by step

					tag, errBool := tagItem.Attr("title") //Tag from site
					errorpack.ErrorBool(errBool)

					if strings.ToLower(tag) == strings.ToLower(sportsTag) { //Single tag image
						accessFlag = true

						historyFile.WriteString("Time of the request: " + currentDate + "\n")
						historyFile.WriteString("News:" + title + "\n")
						historyFile.WriteString(date + time + " Link:" + fullURL + "\n")
						historyFile.WriteString("-----" + "\n")

						fmt.Printf("Time of the request:%s\nNews:%s\n%s %s Link:%s\n-----\n", currentDate, title, date, time, fullURL)
					}
				})
			})
		})
	})
	historyFile.WriteString("\n***\n")
	fmt.Println("Search finished. Press enter to quit.")
	if accessFlag == false {
		fmt.Println("The tag hasn't news related to it. Please, restart the application.") //Information related to tag doesn't exist
	}
}
