package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

// Read the file
func readFile(filePath string) (string, error) {

	// The buffer in bytes package allow for efficient string concatenation
	var fileLines bytes.Buffer

	// use os specific api to open a file
	file, err := os.Open(filePath)

	// file opening requires explicit error checking
	if err != nil {
		return "", err
	}

	// the defer keyword allows an action to be performed contextually to
	// function return
	defer file.Close()

	// The scanner created from io.Reader object defaults to split on line
	fileScanner := bufio.NewScanner(file)

	// Since parameters are passed by value,
	// the append fucntion creates a new slice
	for fileScanner.Scan() {
		fileLines.WriteString(fileScanner.Text())
	}

	return fileLines.String(), fileScanner.Err()
}

// Preprocessing: build the Bad Character table
// For each character of the pattern assign a skip value based on the formula:
// skip = len(pattern) - index(character) - 1
func badCharTable(pattern string) map[rune]int {
	badMap := make(map[rune]int)

	// The cast to rune array allows the counting of actual characters
	patternLength := len([]rune(pattern))

	// For loops on strings iterate over both index and character
	for ind, runeVal := range pattern {
		badMap[runeVal] = patternLength - ind - 1
	}

	return badMap
}

// Main method
func main() {
	filePath := "lorem.txt"
	pattern := "arcu"

	bigText, err := readFile(filePath)
	if err != nil {
		// Fatal prints the error message and stops
		log.Fatal(err)
	}

	badMap := badCharTable(pattern)

	fmt.Println(bigText)
	fmt.Println(badMap)
}
