//Package newstag has function that reads an input tag and returns it
package newstag

import (
	"bufio"
	"fmt"
	"github/sports-parser/errorpack"
	"os"
	"strings"
)

//GetTag reads tag from standart input and returns it
func GetTag() string {
	fmt.Println("Enter Sports.ru news tag (in Russian):")
	sportsTag, err := bufio.NewReader(os.Stdin).ReadString('\n')
	errorpack.ErrorErr(err)
	fmt.Println([]rune(sportsTag))
	return strings.TrimRight(sportsTag, "\r\n")
}
