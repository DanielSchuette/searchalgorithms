package main

import (
	"fmt"
	"time"

	"github.com/DanielSchuette/searchalgorithms"
)

func main() {
	// create a long, unsorted slice
	fmt.Println("creating unsorted slice...")
	unsorted := []int{7, 3, 6, 8, 5, 9, 1, 4, 2, 11, 12, 41, 88, 124, 49, 64, 25, 76, 81,
		14, 3, 41, 19, 12, 91, 42, 62, 99, 12, 52, 19, 57, 34, 13, 63, 71, 34, 96, 17, 63,
		24, 61, 76, 6, 47, 91, 101, 46, 21, 36, 26, 86, 47, 71, 34, 46, 13, 19, 33, 12, 90,
		98, 27, 57, 74, 18, 34, 13, 54, 68, 13, 96, 34, 76, 17, 39, 34, 10, 1, 64, 61, 94,
	}
	for i := 0; i < 12; i++ {
		unsorted = append(unsorted, unsorted...)
	}
	fmt.Printf("slice of len [%d] created\n", len(unsorted))

	// bubble sort, example:
	fmt.Println("starting bubble sort...")
	s1 := time.Now()
	_, c1 := searchalgorithms.BubbleSort(unsorted)
	e1 := time.Now()
	deltaBub := e1.Sub(s1)
	fmt.Printf("bubble sort: time elapsed: %v, iterations: %d\n", deltaBub, c1)

	// insertion sort, example:
	fmt.Println("starting insertion sort...")
	s2 := time.Now()
	_, c2 := searchalgorithms.InsertionSort(unsorted)
	e2 := time.Now()
	deltaIns := e2.Sub(s2)
	fmt.Printf("insertion sort: time elapsed: %v, iterations: %d\n", deltaIns, c2)

	// merge sort, example:
	fmt.Println("starting merge sort...")
	s3 := time.Now()
	_ = searchalgorithms.MergeSort(unsorted)
	e3 := time.Now()
	deltaMer := e3.Sub(s3)
	fmt.Printf("merge sort: time elapsed: %v\n", deltaMer)
}
