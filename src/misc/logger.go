package misc

import (
	"log"
	"os"

	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func init() {
	var loggerCore *zap.Logger
	var err error

	// setup logging
	env := os.Getenv("ENV")
	if env == "production" {
		loggerCore, err = zap.NewProduction()
	} else {
		loggerCore, err = zap.NewDevelopment()
	}

	if err != nil {
		log.Printf("could not setup logger")
		log.Fatalf("%v", err)
	}
	logger = loggerCore.Sugar()
}

func Logger() *zap.SugaredLogger {
	return logger
}
