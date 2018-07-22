package main

import "fmt"

func main() {
	unsorted := []int{7, 3, 6, 8, 5, 9, 1, 4, 2}
	sorted := bubbleSort(unsorted)
	fmt.Println("unsorted:", unsorted)
	fmt.Println("sorted:", sorted)
}

func bubbleSort(in []int) []int {
	ss := make([]int, len(in))
	copy(ss, in)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < (len(ss) - 1); i++ {
			if ss[i] > ss[i+1] {
				ss[i], ss[i+1] = ss[i+1], ss[i]
				swapped = true
			}
		}
	}
	return ss
}
