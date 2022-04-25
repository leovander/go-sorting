package collections

import "github.com/leovander/go-sorting/sorts"

func SortEvens(collection []int) []int {
	evens := evens(collection)
	return sorts.QuickSort(evens, 0, len(evens)-1)
}

func evens(collection []int) []int {
	if len(collection) == 0 {
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
