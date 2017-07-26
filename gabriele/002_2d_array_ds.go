// https://www.hackerrank.com/challenges/2d-array

package main

import "fmt"

func main() {
	n := 6

	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &a[i][j])
		}
	}

	res := -int(^uint(0)>>1) - 1 // min int value

	for i := 0; i < n-2; i++ {
		for j := 0; j < n-2; j++ {
			hglass := a[i][j] + a[i][j+1] + a[i][j+2] +
				a[i+1][j+1] + a[i+2][j] + a[i+2][j+1] + a[i+2][j+2]
			if hglass > res {
				res = hglass
			}
		}
	}

	fmt.Println(res)
}
