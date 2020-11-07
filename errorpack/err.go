package errorpack

import (
	"log"
	"os"
)

type errorList struct {
	errBool string
	errErr  map[int]string
	errSrt  string
}

type errBool struct{}

type otherErr struct{}

func (e *errBool) typeHandler(logFile *log.Logger, err interface{}, listOfErrors errorList, file *os.File) {
	logFile.Println(listOfErrors.errBool)
	file.Close()
	log.Fatal(listOfErrors.errBool)
}

func (o *otherErr) typeHandler(logFile *log.Logger, err interface{}, num int, listOfErrors errorList, file *os.File) {
	logFile.Println(listOfErrors.errErr[num])
	file.Close()
	log.Fatal(listOfErrors.errErr[num])
}

type funcTypes interface {
	typeHandler()
}

//ErrorHandler ...
func ErrorHandler(logFile *log.Logger, err interface{}, num int, file *os.File) {
	listOfErrors := errorsCreating()

	e := &errBool{}
	o := &otherErr{}

	switch err.(type) {
	case bool:
		e.typeHandler(logFile, err, listOfErrors, file)
	default:
		o.typeHandler(logFile, err, num, listOfErrors, file)
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
