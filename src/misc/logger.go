package misc

import (
	"log"

	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func init() {
	// setup logging
	loggerCore, err := zap.NewDevelopment()
	if err != nil {
		log.Printf("could not setup logger")
		log.Fatalf("%v", err)
	}
	logger = loggerCore.Sugar()
}

func Logger() *zap.SugaredLogger {
	return logger
}
