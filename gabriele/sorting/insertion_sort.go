/*
 * Insertion Sort algorithm.
 *
 * Input: sequence of integers A.
 * Output: sorted sequence.
 *
 * At each step j (starting from 1), get the j-th element in A, A[j].
 * We assume that A[0: j - 1] is sorted, which is surely true when j = 1.
 * Go backwards from i = j - 1 to 0 to find the right place in which insert
 * A[j]. When the right place is found, insert A[j] and go to the next step.
 */

package main

import "fmt"

func main() {
  a := [6]int{5, 2, 4, 6, 1, 3}  // average case
  b := [6]int{1, 2, 3, 4, 5, 6}  // best case
  c := [6]int{6, 5, 4, 3, 2, 1}  // worst case
  d := []int{}  // empty array

  insertionSort(a[:])
  insertionSort(b[:])
  insertionSort(c[:])
  insertionSort(d[:])

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
}

func insertionSort(arr []int) {
  for j := 1; j < len(arr); j++ {
    key := arr[j]

    i := j - 1
    for ; i >= 0 && arr[i] > key; i-- {
      arr[i + 1] = arr[i]
    }
    arr[i + 1] = key
  }
}
