// https://www.hackerrank.com/challenges/making-anagrams

package main

import "fmt"

// buildIndex builds map of occurrences from string s
// key: char, val: number of occurrences of char in s
func buildIndex(s string) map[string]int {
	m := make(map[string]int)
	for i := 0; i < len(s); i++ {
		m[string(s[i])] += 1
	}

	return m
}

// checkDifferent returns the number of characters that are
// in str but not in otherIndex
func checkDifferent(str string, otherIndex map[string]int) int {
	different := 0
	for i := 0; i < len(str); i++ {
		s := string(str[i])
		v, ok := otherIndex[s]
		if ok && otherIndex[s] > 0 {
			otherIndex[s] = v - 1
		} else {
			different++
		}
	}

	return different
}

func main() {
	var s1, s2 string
	fmt.Scanf("%s", &s1)
	fmt.Scanf("%s", &s2)

	m1 := buildIndex(s1)
	m2 := buildIndex(s2)

	res := 0
	res += checkDifferent(s1, m2)
	res += checkDifferent(s2, m1)

	fmt.Println(res)
}
