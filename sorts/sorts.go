// Package sorts provides different sorting algorithms that can be performed
// on a slice of ints
package sorts

// Performs recursive the QuickSort in place
func QuickSort(input []int, start int, end int) {
	if start < end {
		var q int = partition(input, start, end)
		QuickSort(input, start, q-1)
		QuickSort(input, q+1, end)
	}
}

func partition(input []int, start int, end int) int {
	var x int = input[end]
	var i int = start - 1

	for j := start; j < end; j++ {
		if input[j] <= x {
			i++
			exchange(input, i, j)
		}
	}

	exchange(input, i+1, end)

	return i + 1
}

func exchange(input []int, a int, b int) {
	var swap int = input[a]
	input[a] = input[b]
	input[b] = swap
}
