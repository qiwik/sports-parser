//Package newstag contains of the function that reads an input tag and returns it,
//and also the function that checks language in the input tag
package newstag

import (
	"bufio"
	"fmt"
	"github/sports-parser/logger"
	"log"
	"os"
	"strings"
	"time"
	"unicode"
)

//GetTag reads tag from standart input and returns it
func GetTag(logFile *os.File) []string {
	fmt.Println("Enter Sports.ru news tag (in Russian) separated by commas:")
	sportsTag, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		err := "Error: reading from the input can not be performed"
		logger.LogErrorClose(logFile, err)
		log.Fatal(err)
	}
	arrayOfTags := strings.Split(sportsTag, ",")

	TagCheck(arrayOfTags, logFile)

	return arrayOfTags
}

//TagCheck verifies language of the input tag
func TagCheck(arrayOfTags []string, logFile *os.File) {
	var checkInput [][]rune
	for i := range arrayOfTags {
		if arrayOfTags[i][0] == 32 {
			arrayOfTags[i] = strings.TrimLeft(arrayOfTags[i], " ")
		}
		arrayOfTags[i] = strings.TrimRight(arrayOfTags[i], "\r\n")
		strToRune := []rune(arrayOfTags[i])
		checkInput = append(checkInput, strToRune)
	}

	for j := range checkInput {
		for k := range checkInput[j] {
			if unicode.Is(unicode.Cyrillic, checkInput[j][k]) != true && checkInput[j][k] != 32 {
				currentTime := time.Now().Format("01-02-2006 15:04:05")
				logFile.WriteString("Invalid tag: not in Russian or more than one space after comma - " + currentTime + "\n------\n")
				logFile.Close()
				log.Fatal("Error: tag not in Russian or more than one space after comma")
			}
		}
	}
}
