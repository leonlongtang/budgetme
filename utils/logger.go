package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func InitLogger() {
	// Set the output to stdout and format to JSON for better parsing in production
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.JSONFormatter{})

	// Set the default logging level to Info
	log.SetLevel(logrus.InfoLevel)
}

// Use this function to log errors, warnings, and info
func GetLogger() *logrus.Logger {
	return log
}
