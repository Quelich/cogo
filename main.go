package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var threadCount int = 4
var complementWord string = "xxx"

func main() {

	fileName := "test1.txt"
	//keyword := "love"

	// startTime := time.Now()
	// fcontent := readFileByLines(fileName)
	// matchedIndices := findMatchedLinesIndex(fcontent, keyword)
	// elapsedTime := time.Since(startTime)

	startTime := time.Now()
	words := readFileByWords(fileName)
	slices := makeSlices(words)
	fmt.Println("---------------------")
	for i := 0; i < len(slices); i++ {

		fmt.Println(slices[i])
		fmt.Println("---------------------")
	}
	elapsedTime := time.Since(startTime)

	// print results
	fmt.Println("Matched Indices:")
	//fmt.Println(matchedIndices)
	fmt.Printf("Elapsed Time:%d\n ns", elapsedTime.Nanoseconds())
}

// Data segmentation for threads
func makeSlices(words []string) [][]string {
	wordsLen := len(words)
	fmt.Printf("Words Length:%d\n", wordsLen)
	startIdx := 0

	/*
	   If the words lengh is odd,
	   add one more word to make it even
	   for convenience in calculations
	*/

	if wordsLen%2 != 0 {
		words = append(words, complementWord)
		wordsLen += 1
	}

	offset := wordsLen / threadCount
	endIndx := offset
	var slices [][]string

	for i := 0; i < threadCount; i++ {
		fmt.Printf("Start: %d - Offset:%d\n", startIdx, endIndx)
		slice := sliceWords(words, startIdx, endIndx)
		slices = append(slices, slice)
		startIdx += offset
		endIndx += offset
	}
	return slices
}

// Read file content line by line
func readFileByLines(fname string) []string {

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

func readFileByWords(fname string) []string {

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

func sliceWords(words []string, from int, to int) []string {
	var slice []string
	for i := from; i < to; i++ {
		word := words[i]
		slice = append(slice, word)
	}
	return slice
}

// Find the indices of matched lines
func findMatchedLinesIndex(keyword string, lines []string, from int, to int) []int {
	var lineIndices []int
	for i := from; i < to; i++ {
		line := lines[i]
		if strings.Contains(line, keyword) {
			lineIndices = append(lineIndices, i+1) // line indices start from one
		}
	}
	return lineIndices
}
