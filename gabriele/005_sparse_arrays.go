// https://www.hackerrank.com/challenges/sparse-arrays

package main

import "fmt"

func main() {
	var n, q int
	var s string
	index := make(map[string]int)

	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &s)
		v, ok := index[s]
		if ok {
			index[s] = v + 1
		} else {
			index[s] = 1
		}
	}

	fmt.Scanf("%d", &q)
	for i := 0; i < q; i++ {
		fmt.Scanf("%s", &s)
		fmt.Println(index[s])
	}
}
