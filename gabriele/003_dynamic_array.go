// https://www.hackerrank.com/challenges/dynamic-array

package main

import "fmt"

func main() {
	var N, Q int
	fmt.Scanf("%d", &N)
	fmt.Scanf("%d", &Q)

	seqList := make([][]int, N)
	for i := 0; i < N; i++ {
		seqList[i] = make([]int, 0)
	}

	lastAnswer := 0

	var q1, q2, q3 int
	for i := 0; i < Q; i++ {
		fmt.Scanf("%d", &q1)
		fmt.Scanf("%d", &q2)
		fmt.Scanf("%d", &q3)

		idx := (q2 ^ lastAnswer) % N

		if q1 == 1 {
			seqList[idx] = append(seqList[idx], q3)
		} else {
			lastAnswer = seqList[idx][q3%len(seqList[idx])]
			fmt.Println(lastAnswer)
		}
	}
}
