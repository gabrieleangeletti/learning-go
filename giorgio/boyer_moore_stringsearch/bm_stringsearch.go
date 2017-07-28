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

	// The buffer in bytes package allow for very efficient string concatenation
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

	// The scanner created from a io.Reader object defaults to split on lines
	fileScanner := bufio.NewScanner(file)

	// Since parameters are passed by value, the append fucntion creates a new slice
	for fileScanner.Scan() {
		fileLines.WriteString(fileScanner.Text())
	}

	return fileLines.String(), fileScanner.Err()
}

func main() {
	filePath := "lorem.txt"

	bigText, err := readFile(filePath)
	if err != nil {
		// Fatal prints the error message and stops
		log.Fatal(err)
	}

	fmt.Println(bigText)
}
