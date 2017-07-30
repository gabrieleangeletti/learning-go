// Go implementation of the Boyer Moore string search algorithm
// https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore_string_search_algorithm

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
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

	return strings.ToLower(fileLines.String()), fileScanner.Err()
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

	// Forces the value of the last rune of the pattern to the pattern length
	badMap[([]rune(pattern))[patternLength-1]] = patternLength

	return badMap
}

// Find occurrencie(s of pattern in text using only the Bad Character rule
// The allignemnt is from right to left --> t_cursor
// The character comparison is from left ot right --> p_cursor
func findMatches(patternString string, textString string, badMap map[rune]int) int {
	occurrencies := 0

	// Convert pattern and text strings to rune slices for efficient access
	pattern := []rune(patternString)
	text := []rune(textString)

	p_len := len(pattern)
	p_cursor := p_len - 1
	t_cursor := 0

	// The while loop is obtained through the for statement
	for t_cursor+p_len <= len(text) {

		t_r := text[t_cursor+p_cursor]
		p_r := pattern[p_cursor]
		fmt.Printf("%c - %c | ", t_r, p_r)

		if p_r == t_r {
			p_cursor--

			if p_cursor < 0 {
				p_cursor = p_len - 1
				t_cursor += p_len
				occurrencies++
				fmt.Printf("String found\n")
			}
		} else {
			p_cursor = p_len - 1
			var skip int

			// If text[i] is in badMap, skip] will get the mapped value and bad will be true
			if val, bad := badMap[t_r]; bad {
				skip = val
			} else {
				skip = p_len
			}
			fmt.Printf("%d \n", skip)
			t_cursor += skip
		}

	}

	return occurrencies
}

// Main method
func main() {
	filePath := "lorem.txt"
	pattern := strings.ToLower("cursus")

	bigText, err := readFile(filePath)
	if err != nil {
		// Fatal prints the error message and stops
		log.Fatal(err)
	}
	fmt.Println(bigText)

	badMap := badCharTable(pattern)
	fmt.Println(badMap)

	occurrencies := findMatches(pattern, bigText, badMap)
	fmt.Println(occurrencies)
}
