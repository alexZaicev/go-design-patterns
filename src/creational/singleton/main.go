package main

import "fmt"

func basicLoggerTest() {
	log := getBasicLoggerInstance()

	log.SetLogLevel(1)
	log.Log("This is a message")

	log = getBasicLoggerInstance()
	log.SetLogLevel(2)
	log.Log("This is a message")

	log = getBasicLoggerInstance()
	log.SetLogLevel(3)
	log.Log("This is a message")
}

func threadSafeLoggerTest() {
	for i := 0; i < 10; i++ {
		go getSafeLoggerInstance()
	}
}

func main() {
	fmt.Println("############################")
	fmt.Println("Basic Logger")
	basicLoggerTest()
	fmt.Println("############################")
	fmt.Println("Thread Safe Logger")
	threadSafeLoggerTest()
}
