package searching_test

import (
	"slices"
	"testing"

	"github.com/Kidpech-code/algorithms-made-clear/searching"
)

func TestLinearSearch(t *testing.T) {
	cases := []struct {
		name   string
		values []int
		target int
		want   int
	}{
		{"empty", nil, 3, -1},
		{"story example", []int{5, 1, 4, 2, 8, 3}, 3, 5},
		{"first item", []int{5, 1, 4, 2, 8, 3}, 5, 0},
		{"first duplicate", []int{4, 2, 4, 2}, 2, 1},
		{"missing", []int{5, 1, 4, 2, 8, 3}, 6, -1},
		{"negative", []int{4, -3, 1}, -3, 1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			before := slices.Clone(tc.values)
			if got := searching.LinearSearch(tc.values, tc.target); got != tc.want {
				t.Fatalf("LinearSearch(%v, %d) = %d; want %d", tc.values, tc.target, got, tc.want)
			}
			if !slices.Equal(tc.values, before) {
				t.Fatal("LinearSearch changed its input")
			}
		})
	}
}

func TestBinarySearch(t *testing.T) {
	cases := []struct {
		name   string
		values []int // Already sorted in ascending order.
		target int
		want   int
	}{
		{"empty", nil, 3, -1},
		{"story example", []int{1, 2, 3, 4, 5, 8}, 3, 2},
		{"first item", []int{1, 2, 3, 4, 5, 8}, 1, 0},
		{"last item", []int{1, 2, 3, 4, 5, 8}, 8, 5},
		{"first duplicate", []int{1, 2, 2, 2, 5}, 2, 1},
		{"missing between", []int{1, 2, 3, 4, 5, 8}, 6, -1},
		{"missing below", []int{1, 2, 3}, 0, -1},
		{"missing above", []int{1, 2, 3}, 4, -1},
		{"one matching", []int{3}, 3, 0},
		{"one missing", []int{3}, 2, -1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			before := slices.Clone(tc.values)
			if got := searching.BinarySearch(tc.values, tc.target); got != tc.want {
				t.Fatalf("BinarySearch(%v, %d) = %d; want %d", tc.values, tc.target, got, tc.want)
			}
			if !slices.Equal(tc.values, before) {
				t.Fatal("BinarySearch changed its input")
			}
		})
	}
}
