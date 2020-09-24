//Package newstag contains of the function that reads an input tag and returns it,
//and also the function that checks language in the input tag
package newstag

import (
	"bufio"
	"fmt"
	"github/sports-parser/errorpack"
	"log"
	"os"
	"strings"
	"time"
	"unicode"
)

//GetTag reads tag from standart input and returns it
func GetTag(logFile *os.File) string {
	fmt.Println("Enter Sports.ru news tag (in Russian):")
	sportsTag, err := bufio.NewReader(os.Stdin).ReadString('\n')
	errorpack.ErrorErr(err, logFile)
	sportsTag = strings.TrimRight(sportsTag, "\r\n")

	TagCheck(sportsTag, logFile)

	return sportsTag
}

//TagCheck verifies language of the input tag
func TagCheck(sportsTag string, logFile *os.File) {
	checkInput := []rune(sportsTag)
	for i := range checkInput {
		if unicode.Is(unicode.Cyrillic, checkInput[i]) != true {
			currentTime := time.Now().Format("01-02-2006 15:04:05")
			logFile.WriteString("Error with the application - " + currentTime)
			logFile.Close()
			log.Fatal("Tag not in Russian")
		}
	}
}
