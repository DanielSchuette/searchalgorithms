package main

import (
	"fmt"

	"github.com/DanielSchuette/searchalgorithms"
)

func main() {
	// bubble sort, example:
	unsorted := []int{7, 3, 6, 8, 5, 9, 1, 4, 2}
	sorted := searchalgorithms.BubbleSort(unsorted)
	fmt.Println("unsorted:", unsorted)
	fmt.Println("sorted:", sorted)

	// insertion sort, example:
}
