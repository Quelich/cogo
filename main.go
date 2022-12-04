package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

var threadCount int = 4
var complementWord string = " "
var matchedIndices []int

func main() {

	fileName := "test1.txt"
	keyword := "love"

	lines := readFileByLines(fileName)
	// CONCURRENT OPERATION
	concurrentOp(lines, keyword)
	// Calculate elapsed time
	startTime := time.Now()

	elapsedTime := time.Since(startTime)

	// print results
	fmt.Printf("Elapsed Time:%d ns\n ", elapsedTime.Nanoseconds())
	fmt.Println(matchedIndices)
}

func concurrentOp(lines []string, keyword string) {
	var wg sync.WaitGroup /* SYNCHRONIZE goroutines */

	linesLen := len(lines)

	// Make the line length even
	if linesLen%2 != 0 {
		lines = append(lines, complementWord)
		linesLen += 1
	}

	offset := linesLen / threadCount
	startIdx := 0
	endIndx := offset

	for i := 0; i < threadCount; i++ {
		fmt.Printf("StartIdx: %d - EndIdx: %d\n", startIdx, endIndx)
		wg.Add(1)
		go func() {
			defer wg.Done()
			getMatchedLinesIndices(keyword, lines, startIdx, endIndx)

		}()
		startIdx += offset
		endIndx += offset
	}
	wg.Wait()
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

// Find the indices of matched lines
func getMatchedLinesIndices(keyword string, lines []string, from int, to int) []int {
	var lineIndices []int
	for i := from; i < to; i++ {
		line := lines[i]
		if strings.Contains(line, keyword) {
			lineIndices = append(lineIndices, i+1) // line indices start from one
		}
	}
	fmt.Println(lineIndices)
	matchedIndices = append(matchedIndices, lineIndices...)
	return lineIndices
}
