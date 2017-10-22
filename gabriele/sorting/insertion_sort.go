package main

import "fmt"

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
