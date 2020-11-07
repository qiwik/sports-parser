//Package newstag contains of the function that reads an input tag and returns it,
//and also the function that checks language in the input tag
package newstag

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sports-parser/errorpack"
	"strings"
	"unicode"
)

//GetTag reads tag from standart input and returns it
func GetTag(logFile *log.Logger, file *os.File) []string {
	var err interface{}
	fmt.Println("Enter Sports.ru news tag (in Russian) separated by commas:")
	sportsTag, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		num := 5
		errorpack.ErrorHandler(logFile, err, num, file)
	}
	arrayOfTags := strings.Split(sportsTag, ",")

	tagCheck(arrayOfTags, logFile, file)

	return arrayOfTags
}

//TagCheck verifies language of the input tag
func tagCheck(arrayOfTags []string, logFile *log.Logger, file *os.File) {
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
				var err interface{} = "error"
				num := 4
				errorpack.ErrorHandler(logFile, err, num, file)
			}
		}
	}
}
