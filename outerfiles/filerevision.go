package outerfiles

import (
	"github/sports-parser/errorpack"
	"io/ioutil"
)

//Revision helps to find history.txt in the root
func Revision() bool {
	fileFromDir, err := ioutil.ReadDir(".")
	errorpack.ErrorErr(err)

	for _, file := range fileFromDir {
		if file.Name() == "history.txt" {
			return true
		}
	}
	return false
}
