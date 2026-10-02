package sorting_test

import (
	"slices"
	"testing"

	"github.com/Kidpech-code/algorithms-made-clear/sorting"
)

func TestSorts(t *testing.T) {
	sorts := []struct {
		name string
		sort func([]int)
	}{
		{"Bubble", sorting.BubbleSort},
		{"Merge", sorting.MergeSort},
		{"Quick", sorting.QuickSort},
	}
	cases := []struct {
		name  string
		input []int
		want  []int
	}{
		{"empty", nil, nil},
		{"one number", []int{7}, []int{7}},
		{"story example", []int{5, 1, 4, 2, 8, 3}, []int{1, 2, 3, 4, 5, 8}},
		{"duplicates", []int{4, 2, 4, 1, 2}, []int{1, 2, 2, 4, 4}},
		{"negatives", []int{0, -3, 8, -1}, []int{-3, -1, 0, 8}},
		{"already sorted", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{"descending", []int{4, 3, 2, 1}, []int{1, 2, 3, 4}},
	}

	for _, algorithm := range sorts {
		t.Run(algorithm.name, func(t *testing.T) {
			for _, tc := range cases {
				t.Run(tc.name, func(t *testing.T) {
					got := slices.Clone(tc.input)
					algorithm.sort(got)
					if !slices.Equal(got, tc.want) {
						t.Fatalf("sort(%v) = %v; want %v", tc.input, got, tc.want)
					}
				})
			}
		})
	}
}
