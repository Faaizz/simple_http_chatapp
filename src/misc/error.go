package misc

import (
	"os"
)

// ErrorCheck prints the err and msg to the logger if err != nil.
func ErrorCheck(err interface{}, msg string) {
	if err != nil {
		logger.Errorln(err)
		logger.Infoln(msg)
	}
}

// ErrorCheckExit calls ErrorCheck with the (err, msg) pair and exit the program with a non-zero return code if err != nil.
func ErrorCheckExit(err interface{}, msg string) {
	if err != nil {
		ErrorCheck(err, msg)
		os.Exit(1)
	}
}
