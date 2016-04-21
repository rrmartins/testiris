package main

import (
	"fmt"
	"log"
	"os"
)

func simpleLogging() {
	fmt.Println("---------- simple logging ------")
	log.Println("hello world")
	log.Println("this is a simple error")
}

func formattingLogging() {
	fmt.Println("-------------- formattingLogging ----------")
	var warning *log.Logger

	warning = log.New(
		os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	warning.Println("This is warning message 1")
	warning.Println("This is warning message 2")
}

func fileLogging() {
	fmt.Println("-------------- file Logging ----------")
	file, err := os.OpenFile("./myapp.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Failed to open log file")
		return
	}

	var logFile *log.Logger
	logFile = log.New(file,
		"APP: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	logFile.Println("This is error message 1")
	logFile.Println("This is error message 2")
	logFile.Println("This is error message 3")
	fmt.Println("Done")
}

func main() {
	// simpleLogging()
	// formattingLogging()
	fileLogging()
}
