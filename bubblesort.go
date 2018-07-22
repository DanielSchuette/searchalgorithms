package sortingalgorithms

func bubbleSort(in []int) []int {
	ss := make([]int, len(in))
	copy(ss, in)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < (len(ss) - 1); i++ {
			if ss[i] > ss[i+1] {
				ss[i], ss[i+1] = ss[i+1], ss[i]
				swapped = true
			}
		}
	}
	return ss
}
