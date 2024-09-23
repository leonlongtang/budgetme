package utils

import (
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	log     *logrus.Logger
	logOnce sync.Once
)

func InitLogger() {
	log = logrus.New()

	// Set the output to stdout and format to JSON for better parsing in production
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	log.SetLevel(logrus.DebugLevel)

}

// Use this function to log errors, warnings, and info
func GetLogger() *logrus.Logger {
	logOnce.Do(InitLogger)
	return log
}
