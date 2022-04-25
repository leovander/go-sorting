package main

import (
	"fmt"

	"github.com/leovander/go-sorting/collections"
)

func main() {
	input := []int{2, 8, 7, 1, 3, 5, 6, 4}
	fmt.Println("Input", input)

	var sorted []int = collections.SortEvens(input)
	fmt.Println("Sorted", sorted)
}
