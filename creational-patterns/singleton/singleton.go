package main

import (
	"fmt"
	"sync"
)

type Logger struct {
	logLevel int
}

func (l *Logger) Log(s string) {
	fmt.Printf("%d: %s\n", l.logLevel, s)
}

func (l *Logger) SetLogLevel(level int) {
	l.logLevel = level
}

var once sync.Once
var logger *Logger

func getLoggerInstance() *Logger {
	once.Do(func() {
		fmt.Println("Creating a new logger instance")
		logger = &Logger{logLevel: 1}
	})

	fmt.Println("Returning logger instance")
	return logger
}
