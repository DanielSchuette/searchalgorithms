package searchalgorithms

// BubbleSort sorts the input with time complexity 0(n^2); returns a new, sorted slice
// it does not modify the input slice because it internally makes a copy as well as an
// integer value indicating how many iterations were necessary to sort the input slice
func BubbleSort(in []int) ([]int, int) {
	var sortCount int
	ss := make([]int, len(in))
	copy(ss, in)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < (len(ss) - 1); i++ {
			if ss[i] > ss[i+1] {
				ss[i], ss[i+1] = ss[i+1], ss[i]
				swapped = true
				sortCount++
			}
		}
	}
	return ss, sortCount
}
