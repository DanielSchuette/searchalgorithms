package main

import (
	"fmt"
	"time"

	"github.com/DanielSchuette/searchalgorithms"
)

func main() {
	// create an unsorted array and print it
	unsorted := []int{7, 3, 6, 8, 5, 9, 1, 4, 2, 11,
		14, 3, 41, 19, 12, 91, 42, 62, 99, 12,
		24, 61, 76, 6, 47, 91, 101, 46, 21, 36}
	fmt.Printf("unsorted: %v\n", unsorted)

	// bubble sort, example:
	start := time.Now()
	sortedBub, c := searchalgorithms.BubbleSort(unsorted)
	end := time.Now()
	deltaBub := end.Sub(start)
	fmt.Printf("bubble sorted: %v, time elapsed: %v, iterations: %d\n", sortedBub, deltaBub, c)

	// insertion sort, example:
	start = time.Now()
	sortedIns, c := searchalgorithms.InsertionSort(unsorted)
	end = time.Now()
	deltaIns := end.Sub(start)
	fmt.Printf("insertion sorted: %v, time elapsed: %v, iterations: %d\n", sortedIns, deltaIns, c)

	// merge sort, example:
	start = time.Now()
	sortedMer := searchalgorithms.MergeSort(unsorted)
	end = time.Now()
	deltaMer := end.Sub(start)
	fmt.Printf("merge sorted: %v, time elapsed: %v\n", sortedMer, deltaMer)
}
