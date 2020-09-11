//Package errorpack cheks error state for bool or error variable
package errorpack

import "log"

//ErrorErr checks error state if variable is error, not bool
func ErrorErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//ErrorBool checks error state if variable is bool, not error
func ErrorBool(err bool) {
	if err != true {
		log.Fatal(err)
	}
}
