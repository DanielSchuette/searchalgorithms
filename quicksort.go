package searchalgorithms

// QuickSort implements a quick sort algorithm; it takes a slice of ints and returns a sorted slice of ints
// QuickSort operates with time complexity of O(n*logn)
func QuickSort(in []int) []int {
	if len(in) < 2 {
		return in
	}
	pivot := in[len(in)-1]
	less := make([]int, 0)
	more := make([]int, 0)
	for i := 0; i < (len(in) - 1); i++ {
		if in[i] <= pivot {
			less = append(less, in[i])
		} else {
			more = append(more, in[i])
		}
	}
	tmp := append(QuickSort(less), pivot)
	return append(tmp, QuickSort(more)...)
}
