package main

import (
	"fmt"

	"github.com/DanielSchuette/searchalgorithms"
)

func main() {
	// bubble sort, example:
	unsorted := []int{7, 3, 6, 8, 5, 9, 1, 4, 2}
	sortedBub := searchalgorithms.BubbleSort(unsorted)
	fmt.Printf("unsorted: %v\n", unsorted)
	fmt.Printf("bubble sorted: %v\n", sortedBub)

	// insertion sort, example:
	sortedIns := searchalgorithms.InsertionSort(unsorted)
	fmt.Printf("unsorted: %v\n", unsorted)
	fmt.Printf("insertion sorted: %v\n", sortedIns)
}
