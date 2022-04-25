package sorts_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/leovander/go-sorting/sorts"
)

func TestQuickSort(t *testing.T) {
	input := []int{2, 8, 7, 1, 3, 5, 6, 4}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8}
	sorts.QuickSort(input, 0, len(input)-1)
	if diff := cmp.Diff(expected, input); diff != "" {
		t.Errorf(diff)
	}
}

func TestQuickSortEmpty(t *testing.T) {
	input := []int{}
	sorts.QuickSort(input, 0, len(input)-1)
	if len(input) != 0 {
		t.Error("incorrect result: expected empty input, got", input)
	}
}

func TestQuickSortSingle(t *testing.T) {
	input := []int{1}
	expected := []int{1}
	sorts.QuickSort(input, 0, len(input)-1)
	if diff := cmp.Diff(expected, input); diff != "" {
		t.Errorf(diff)
	}
}

func TestQuickSortTuple(t *testing.T) {
	input := []int{2, 1}
	expected := []int{1, 2}
	sorts.QuickSort(input, 0, len(input)-1)
	if diff := cmp.Diff(expected, input); diff != "" {
		t.Errorf(diff)
	}
}
