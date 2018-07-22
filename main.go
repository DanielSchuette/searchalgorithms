package main

import (
	"fmt"

	"github.com/DanielSchuette/sortingalgorithms"
)

func main() {
	// bubble sort, example:
	unsorted := []int{7, 3, 6, 8, 5, 9, 1, 4, 2}
	sorted := sortingalgorithms.bubbleSort(unsorted)
	fmt.Println("unsorted:", unsorted)
	fmt.Println("sorted:", sorted)

	// insertion sort, example:
}
