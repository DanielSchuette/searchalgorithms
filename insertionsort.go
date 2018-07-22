package searchalgorithms

// InsertionSort implements an insertion sort algorithm that takes a slice of integers
// and returns a sorted copy of that slice without altering the input;
// works well on 'almost' sorted slices with O(n^2)
func InsertionSort(in []int) []int {
	ss := make([]int, len(in))
	copy(ss, in)

	// loop over the copied slice
	for i := 0; i < len(ss); i++ {

		// don't do anything with the first element, it can be considered 'sorted' at this point
		if i == 0 {
			continue
		}

		// for all other elements, iterate over all elements before the current element at ss[i]
		// and test if that particular element ss[j] is larger than ss[i]; it that is the case,
		// we found the position to insert the current ss[i] because we sort in ascending order
		// and because the slice is always sorted up to the current ss[i]
		for j := 0; j < i; j++ {
			if ss[i] < ss[j] {

				// save the current element at ss[i]
				tmp := ss[i]

				// shift all elements starting at ss[i-1] backwards to ss[j+1] so that the
				// correct order is keep while making insertion possible
				for k := i; k > j; k-- {
					ss[k] = ss[k-1]
				}

				// insert the value of the current element ss[i] at the right position ss[j]
				ss[j] = tmp
			}
		}
	}
	return ss
}
