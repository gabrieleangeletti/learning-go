// https://www.hackerrank.com/challenges/array-left-rotation

package main

import "fmt"

func main() {
	var N, d int
	fmt.Scanf("%d", &N)
	fmt.Scanf("%d", &d)

	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	res := make([]int, N)
	for i := 0; i < N; i++ {
		res[(i-d+N)%N] = arr[i]
	}

	for i := 0; i < N; i++ {
		fmt.Printf("%d ", res[i])
	}
}
