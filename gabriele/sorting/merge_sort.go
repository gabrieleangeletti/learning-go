package main

import (
	"fmt"
	"math"
)

func main() {
	a := [6]int{5, 2, 4, 6, 1, 3} // average case
	b := [6]int{1, 2, 3, 4, 5, 6} // best case
	c := [6]int{6, 5, 4, 3, 2, 1} // worst case
	d := []int{}                  // empty array

	MergeSort(a[:])
	MergeSort(b[:])
	MergeSort(c[:])
	MergeSort(d[:])

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

// MergeSort ...
//
func MergeSort(arr []int) {
	mergeSortAux(arr, 0, len(arr)-1)
}

func mergeSortAux(arr []int, p int, r int) {
	if p < r {
		q := int(math.Floor(float64((p + r) / 2)))
		mergeSortAux(arr, p, q)
		mergeSortAux(arr, q+1, r)
		merge(arr, p, q, r)
	}
}

func merge(arr []int, p int, q int, r int) {
	n1 := q - p + 1
	n2 := r - q

	L := make([]int, n1+1, n1+1)
	for i := 0; i < n1; i++ {
		L[i] = arr[p+i]
	}

	R := make([]int, n2+1, n2+1)
	for i := 0; i < n2; i++ {
		R[i] = arr[q+i+1]
	}

	L[n1] = math.MaxInt32
	R[n2] = math.MaxInt32

	i, j := 0, 0

	for k := p; k < r+1; k++ {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
	}
}
