package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
)

func main() {
	fmt.Println("Hello, World!")

	//keyword := "love";
	fileName := "test.txt"
	fcontent := readFile(fileName)
	printFileRanged(fcontent, 0, 10)
	// read file
	// find the keywords
	// list the keywords
	// print output
}

func readFile(fname string) []string {
	
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal("Error opening file", err)
	}

	Scanner := bufio.NewScanner(file)
	Scanner.Split(bufio.ScanWords)
	
	var words []string
	
	for Scanner.Scan() {
		words = append(words, Scanner.Text())
	}

	return words
}

func printFile(fcontent []string) {
	
	fmt.Printf("File Content:\n%s", fcontent)
}

func printFileRanged(words []string, from int, to int) {
	for i := from; i < to; i++ {
		word := words[i]
		fmt.Println(word)
	}
}