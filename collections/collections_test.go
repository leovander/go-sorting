package collections

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_evens(t *testing.T) {
	input := []int{1, 2}
	expected := []int{2}
	result := evens(input)
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Errorf(diff)
	}
}

func Test_isEven(t *testing.T) {
	input := 4
	if !isEven(input) {
		t.Error("incorrect result: expected true")
	}
}
