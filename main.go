//TODO: split lines and use threads
package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")

	//keyword := "love";
	fileName := "test.txt"
	fcontent := readFile(fileName)
	keyword := "love"
	//printFileRanged(fcontent, 0, 10)
	matchedIndices := findMatchedLines(fcontent, keyword)
	fmt.Println("Matched Indices:")
	fmt.Println(matchedIndices)
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
	Scanner.Split(bufio.ScanLines)
	
	var lines []string
	
	for Scanner.Scan() {
		lines = append(lines, Scanner.Text())
	}

	return lines
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

func findMatchedLines(lines []string, keyword string) []int {
	
	var lineIndices []int

	for i := 0; i < len(lines) ; i++ {
		line := lines[i]
		if  strings.Contains(line, keyword){
			lineIndices = append(lineIndices, i + 1) // line indices start from one
		}
	}

	return lineIndices
}
