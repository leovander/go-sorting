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

func TestQuickSortLarge(t *testing.T) {
	input := []int{
		48, 79, 67, 44, 31, 64, 6, 46, 34, 46, 30, 86, 67, 46, 40, 85, 93, 51, 62,
		23, 8, 57, 77, 25, 19, 33, 60, 65, 78, 23, 61, 95, 19, 90, 38, 46, 71, 67,
		62, 91, 96, 76, 3, 62, 74, 58, 9, 34, 78, 31, 69, 48, 1, 42, 59, 79, 72, 38,
		10, 77, 73, 20, 70, 48, 31, 62, 8, 10, 82, 26, 32, 67, 13, 19, 14, 55, 2,
		55, 37, 65, 21, 71, 23, 74, 53, 21, 11, 52, 34, 84, 23, 9, 64, 29, 95, 21,
		90, 57, 62, 61}
	expected := []int{1, 2, 3, 6, 8, 8, 9, 9, 10, 10, 11, 13, 14, 19, 19, 19, 20,
		21, 21, 21, 23, 23, 23, 23, 25, 26, 29, 30, 31, 31, 31, 32, 33, 34, 34, 34,
		37, 38, 38, 40, 42, 44, 46, 46, 46, 46, 48, 48, 48, 51, 52, 53, 55, 55, 57,
		57, 58, 59, 60, 61, 61, 62, 62, 62, 62, 62, 64, 64, 65, 65, 67, 67, 67, 67,
		69, 70, 71, 71, 72, 73, 74, 74, 76, 77, 77, 78, 78, 79, 79, 82, 84, 85, 86,
		90, 90, 91, 93, 95, 95, 96}
	sorts.QuickSort(input, 0, len(input)-1)
	if diff := cmp.Diff(expected, input); diff != "" {
		t.Errorf(diff)
	}
}
