package main

import (
	"fmt"
	"time"

	"github.com/DanielSchuette/searchalgorithms"
)

func main() {
	// create a long, unsorted slice
	fmt.Println("creating unsorted slice...")
	unsorted := []int{7, 3, 6, 8, 5, 9, 1, 4, 2, 11, 12, 41, 88, 124, 49, 64, 25, 76, 81}
	for i := 0; i < 15; i++ {
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

	// quick sort, example:
	fmt.Println("starting quick sort...")
	s4 := time.Now()
	_ = searchalgorithms.QuickSort(unsorted)
	e4 := time.Now()
	deltaQui := e4.Sub(s4)
	fmt.Printf("quick sort: time elapsed: %v\n", deltaQui)
}
