package recursion_test

import (
	"slices"
	"testing"

	"github.com/Kidpech-code/algorithms-made-clear/recursion"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   int
	}{
		{name: "nil", want: 0},
		{name: "empty", values: []int{}, want: 0},
		{name: "one value", values: []int{7}, want: 7},
		{name: "chapter example", values: []int{5, 1, 4, 2, 8, 3}, want: 23},
		{name: "negative values", values: []int{-5, 2, -1}, want: -4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			before := slices.Clone(tt.values)
			if got := recursion.Sum(tt.values); got != tt.want {
				t.Fatalf("Sum(%v) = %d; want %d", tt.values, got, tt.want)
			}
			if !slices.Equal(tt.values, before) {
				t.Fatalf("Sum changed input: got %v; want %v", tt.values, before)
			}
		})
	}
}
