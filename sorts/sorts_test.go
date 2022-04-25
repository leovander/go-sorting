package sorts

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_partition(t *testing.T) {
	input := []int{2, 8, 7, 1, 3, 5, 6, 4}
	result := partition(input, 0, len(input)-1)
	if result != 3 {
		t.Error("incorrect result: expected 5, got", result)
	}
}

func Test_exchange(t *testing.T) {
	input := []int{1, 2}
	expected := []int{2, 1}
	exchange(input, 0, 1)
	if diff := cmp.Diff(expected, input); diff != "" {
		t.Errorf(diff)
	}
}
