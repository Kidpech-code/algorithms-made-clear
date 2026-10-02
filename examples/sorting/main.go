package main

import (
	"fmt"
	"slices"

	"github.com/Kidpech-code/algorithms-made-clear/sorting"
)

func main() {
	original := []int{5, 1, 4, 2, 8, 3}
	fmt.Println("Before:", original)

	algorithms := []struct {
		name string
		sort func([]int)
	}{
		{"Bubble", sorting.BubbleSort},
		{"Merge ", sorting.MergeSort},
		{"Quick ", sorting.QuickSort},
	}
	for _, algorithm := range algorithms {
		values := slices.Clone(original)
		algorithm.sort(values)
		fmt.Println(algorithm.name+":", values)
	}
}
