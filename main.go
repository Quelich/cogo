package main

import (
	"fmt"
	"log"
	"os"
)


func main() {
    fmt.Println("Hello, World!")
	
	file, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal("Error opening file", err);
	}
	
	// read file
	fmt.Printf("File Content:\n%s", file)
	// find the keywords
	// list the keywords
	// print output 
}