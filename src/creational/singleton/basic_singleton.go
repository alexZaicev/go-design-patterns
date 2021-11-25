package main

import "fmt"

type BasicLogger struct {
	logLevel int
}

func (l *BasicLogger) Log(s string) {
	fmt.Println(l.logLevel, ":", s)
}

func (l *BasicLogger) SetLogLevel(logLevel int) {
	l.logLevel = logLevel
}

var basicLogger *BasicLogger

func getBasicLoggerInstance() *BasicLogger {
	if basicLogger == nil {
		fmt.Println("Creating basicLogger")
		basicLogger = &BasicLogger{}
	}
	fmt.Println("Returning basicLogger")
	return basicLogger
}
