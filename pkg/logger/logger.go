package logger

import (
	"log"
	"os"
)

type Logger struct {
	*log.Logger
}

func New() *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "[LIBRARY-SERVICE] ", log.LstdFlags|log.Lshortfile),
	}
}

func (l *Logger) Info(msg string) {
	l.Printf("INFO: %s", msg)
}

func (l *Logger) Error(msg string) {
	l.Printf("ERROR: %s", msg)
}

func (l *Logger) Fatal(msg string) {
	l.Fatalf("FATAL: %s", msg)
}