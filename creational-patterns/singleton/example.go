package main

import "fmt"

func main() {
	//logger := getLoggerInstance()
	//logger.SetLogLevel(1)

	//logger.Log("Hello there")
	//logger.SetLogLevel(2)
	//logger.Log("What's up?")

	for i := 1; i < 11; i++ {
		go getLoggerInstance()
	}

	fmt.Scanln()
}
