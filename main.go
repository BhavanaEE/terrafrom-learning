package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func ReadFile() {

	fmt.Println("Reading a file in Go lang")
	fileName := "test.txt"
	
	data, err := ioutil.ReadFile("hello.txt")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nFile Name: %s", fileName)
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s", data)

}

// main function
func main() {

	ReadFile()
}
