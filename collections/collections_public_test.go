package collections_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/leovander/go-sorting/collections"
)

func TestSortEvens(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		input := []int{}
		expected := []int{}
		result := collections.SortEvens(input)
		if diff := cmp.Diff(expected, result); diff != "" {
			t.Errorf(diff)
		}
	})

	t.Run("slice of 8 int", func(t *testing.T) {
		input := []int{2, 8, 7, 1, 3, 5, 6, 4}
		expected := []int{2, 4, 6, 8}
		result := collections.SortEvens(input)
		if diff := cmp.Diff(expected, result); diff != "" {
			t.Errorf(diff)
		}
	})
}
