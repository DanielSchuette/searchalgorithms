package searchalgorithms

import (
	"fmt"
	"log"
)

// MergeSort makes a copy of the input slice and sorts it using a mergesort algorithm
// time complexity of O(n*logn)
func MergeSort(in []int) []int {
	if len(in) < 2 {
		return in
	}
	fmt.Printf("calling merged result of a: %v, b: %v\n", in[:(len(in)/2)], in[(len(in)/2):])
	return merge(MergeSort(in[:(len(in)/2)]), MergeSort(in[(len(in)/2):]))
}

// merge merges to sorted input slices
func merge(sa, sb []int) []int {
	fmt.Printf("merge got a: %v, b: %v\n", sa, sb)
	var aCount int
	var bCount int
	returnSlice := make([]int, 0)
Loop:
	for {
		if aCount == (len(sa) - 1) {
			fmt.Printf("append to end; input: %v, count: %v\n", sa, aCount)
			returnSlice = append(returnSlice, sb[aCount:]...)
			break Loop
		} else if bCount == (len(sb) - 1) {
			fmt.Printf("append to end; input: %v, count: %v\n", sb, bCount)
			returnSlice = append(returnSlice, sa[aCount:]...)
			break Loop
		} else if sa[aCount] < sb[bCount] {
			fmt.Printf("append; input: %v, count: %v\n", sa, aCount)
			returnSlice = append(returnSlice, sa[aCount])
			aCount++
		} else if sa[aCount] > sb[bCount] {
			fmt.Printf("append; input: %v, count: %v\n", sb, bCount)
			returnSlice = append(returnSlice, sb[bCount])
			bCount++

		} else {
			log.Fatal("don't know what to do")
		}
	}
	return returnSlice
}
