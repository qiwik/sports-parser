//Package errorpack cheks error state for bool or error variable
package errorpack

import (
	"log"
	"os"
	"time"
)

//ErrorErr checks error state if variable is error, not bool
func ErrorErr(err error, logFile *os.File) {
	if err != nil {
		currentTime := time.Now().Format("01-02-2006 15:04:05")
		logFile.WriteString("Error with the application - " + currentTime + "\n------\n")
		logFile.Close()
		log.Fatal(err)
	}
}

//ErrorBool checks error state if variable is bool, not error
func ErrorBool(err bool, logFile *os.File) {
	if err != true {
		currentTime := time.Now().Format("01-02-2006 15:04:05")
		logFile.WriteString("Error with the application - " + currentTime + "\n------\n")
		logFile.Close()
		log.Fatal(err)
	}
}
