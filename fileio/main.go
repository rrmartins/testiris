package main

import (
	"fmt"
	"io/ioutil"
)

func writeFile(message string) {
	bytes := []byte(message)
	ioutil.WriteFile("./testgo.txt", bytes, 0644)
	fmt.Println("created a file")
}

func readFile() {
	data, _ := ioutil.ReadFile("./testgo.txt")
	fmt.Println("file content:")
	fmt.Println(string(data))
}

func main() {
	fmt.Println("writing data into a file")
	writeFile("welcome to go")
	readFile()

	fmt.Println("reading data from a file")
	readFile()
}
