package errorpack

import (
	"log"
	"os"
	"time"
)

type errorList struct {
	errBool string
	errErr  map[int]string
	errSrt  string
}

type errBool struct{}

type otherErr struct{}

func (e *errBool) typeHandler(logFile *os.File, err interface{}, listOfErrors errorList) {
	currentTime := time.Now().Format("01-02-2006 15:04:05")
	logFile.WriteString(listOfErrors.errBool + " " + currentTime + "\n------\n")
	logFile.Close()
	log.Fatal(listOfErrors.errBool)
}

func (o *otherErr) typeHandler(logFile *os.File, err interface{}, num int, listOfErrors errorList) {
	currentTime := time.Now().Format("01-02-2006 15:04:05")
	logFile.WriteString(listOfErrors.errErr[num] + " " + currentTime + "\n------\n")
	logFile.Close()
	log.Fatal(listOfErrors.errErr[num])
}

type funcTypes interface {
	typeHandler()
}

func ErrorHandler(logFile *os.File, err interface{}, num int) {
	listOfErrors := errorsCreating()

	e := &errBool{}
	o := &otherErr{}

	switch err.(type) {
	case bool:
		e.typeHandler(logFile, err, listOfErrors)
	default:
		o.typeHandler(logFile, err, num, listOfErrors)
	}
}

func errorsCreating() (listOfErrors errorList) {
	errMap := map[int]string{
		1: "Error: the request failed",
		2: "Error: html reading can not be performed",
		3: "Error: history.txt can not be created or opened",
		4: "Error: tag not in Russian or more than one space after comma",
		5: "Error: reading from the input can not be performed",
	}

	listOfErrors = errorList{errBool: "Error: html reading can not be performed",
		errErr: errMap, errSrt: "Invalid tag: not in Russian or more than one space after comma - ",
	}

	return listOfErrors
}
