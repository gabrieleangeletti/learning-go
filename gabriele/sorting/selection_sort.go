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

	SelectionSort(b[:])
	SelectionSort(c[:])
	SelectionSort(d[:])
	SelectionSort(a[:])

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

// SelectionSort ...
//
// Selection sort algorithm.
// Input: sequence of integers A.
// Output: sorted sequence.
//
// At time step 0, get the smallest element in A and swap it
// with A[0]. At time step 1, get the smallest element in A[1:]
// and swap it with A[1]. Stop at step n - 2, because in step n - 1,
// A[n - 1:] is just one element, which is the largest one in the
// array, and the latest.
//
func SelectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		swap(arr, i, minValue(arr, i, len(arr)))
	}
}

func minValue(arr []int, start int, end int) int {
	minValue := math.MaxInt32
	minIndex := -1
	for i := start; i < end; i++ {
		if arr[i] < minValue {
			minValue = arr[i]
			minIndex = i
		}
	}
	return minIndex
}

func swap(arr []int, index1 int, index2 int) {
	tmp := arr[index1]
	arr[index1] = arr[index2]
	arr[index2] = tmp
}
