package main

import (
	"fmt"
	"sync"
)

type SafeLogger struct {
	logLevel int
}

func (l *SafeLogger) Log(s string) {
	fmt.Println(l.logLevel, ":", s)
}

func (l *SafeLogger) SetLogLevel(logLevel int) {
	l.logLevel = logLevel
}

var safeLogger *SafeLogger

var once sync.Once

func getSafeLoggerInstance() *SafeLogger {
	once.Do(func() {
		fmt.Println("Creating safeLogger")
		safeLogger = &SafeLogger{}
	})
	fmt.Println("Returning safeLogger")
	return safeLogger
}
