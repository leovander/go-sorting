// Package collections provides various utilities to perform on slices and arrays
package collections

import "github.com/leovander/go-sorting/sorts"

// Filters out odd ints from a given slice and returns a sorted slice
func SortEvens(collection []int) []int {
	evens := evens(collection)
	sorts.QuickSort(evens, 0, len(evens)-1)
	return evens
}

// Returns a slice of a slice containing only even ints
func evens(collection []int) []int {
	if len(collection) <= 1 {
		return collection
	}

	idx := 0
	for _, v := range collection {
		if isEven(v) {
			collection[idx] = v
			idx++
		}
	}

	collection = collection[:idx]

	return collection
}

func isEven(value int) bool {
	return value%2 == 0
}
