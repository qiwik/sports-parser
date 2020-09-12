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
	"unicode"
)

//GetTag reads tag from standart input and returns it
func GetTag() string {
	fmt.Println("Enter Sports.ru news tag (in Russian):")
	sportsTag, err := bufio.NewReader(os.Stdin).ReadString('\n')
	errorpack.ErrorErr(err)
	sportsTag = strings.TrimRight(sportsTag, "\r\n")

	TagCheck(sportsTag)

	return sportsTag
}

//TagCheck verifies language of the input tag
func TagCheck(sportsTag string) {
	checkInput := []rune(sportsTag)
	for i := range checkInput {
		if unicode.Is(unicode.Cyrillic, checkInput[i]) != true {
			log.Fatal("Tag not in Russian")
		}
	}
}
