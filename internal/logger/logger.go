package logger

import (
	"log"
	"os"
)

type Logger struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
}

func NewLogger() Logger {
	InfoLogger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	return Logger{
		InfoLogger:  InfoLogger,
		ErrorLogger: ErrorLogger,
	}
}
