package logger

import (
	"log"
	"os"
)

// Logger is a global logger instance
var Logger *log.Logger

// InitLogger initializes the logger with a file
func InitLogger(logFile string) {
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}

	Logger = log.New(file, "", log.LstdFlags|log.Lshortfile)
	log.SetOutput(file)
	log.Println("Logger initialized successfully.")
}
