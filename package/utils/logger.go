package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

func InitLogger() (*logrus.Logger, *logrus.Logger) {
	currentDate := time.Now().Format("2006-01-02")

	if err := os.MkdirAll("logs/error", os.ModePerm); err != nil {
		logrus.Fatalf("Failed to create error log directory: %v", err)
	}
	
	if err := os.MkdirAll("logs/debug", os.ModePerm); err != nil {
		logrus.Fatalf("Failed to create debug log directory: %v", err)
	}

	errorLogFilePath := filepath.Join("logs/error", fmt.Sprintf("error-%s.log", currentDate))
	debugLogFilePath := filepath.Join("logs/debug", fmt.Sprintf("debug-%s.log", currentDate))

	errorLogFile, err := os.OpenFile(errorLogFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		logrus.Fatalf("Failed to open error log file: %v", err)
	}

	debugLogFile, err := os.OpenFile(debugLogFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		logrus.Fatalf("Failed to open debug log file: %v", err)
	}

	// error logger
	errorLogger := logrus.New()
	errorLogger.SetOutput(errorLogFile)
	errorLogger.SetFormatter(&logrus.JSONFormatter{})
	errorLogger.SetLevel(logrus.ErrorLevel)

	// debug logger
	debugLogger := logrus.New()
	debugLogger.SetOutput(debugLogFile)
	debugLogger.SetFormatter(&logrus.JSONFormatter{})
	debugLogger.SetLevel(logrus.DebugLevel)

	return errorLogger, debugLogger
}