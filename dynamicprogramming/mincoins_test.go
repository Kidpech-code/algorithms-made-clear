package dynamicprogramming_test

import (
	"slices"
	"testing"

	"github.com/Kidpech-code/algorithms-made-clear/dynamicprogramming"
)

func TestMinCoins(t *testing.T) {
	implementations := []struct {
		name string
		find func([]int, int) (int, error)
	}{
		{name: "memoization", find: dynamicprogramming.MinCoinsMemo},
		{name: "tabulation", find: dynamicprogramming.MinCoinsTable},
	}
	tests := []struct {
		name    string
		coins   []int
		amount  int
		want    int
		wantErr bool
	}{
		{name: "chapter example", coins: []int{1, 3, 4}, amount: 6, want: 2},
		{name: "zero amount and no coins", amount: 0, want: 0},
		{name: "zero amount", coins: []int{1, 3, 4}, amount: 0, want: 0},
		{name: "impossible amount", coins: []int{4, 6}, amount: 5, want: -1},
		{name: "reachable after impossible smaller amounts", coins: []int{3, 4}, amount: 7, want: 2},
		{name: "no coins", amount: 5, want: -1},
		{name: "reordered coins", coins: []int{4, 1, 3}, amount: 6, want: 2},
		{name: "negative amount", coins: []int{1, 3}, amount: -1, wantErr: true},
		{name: "zero coin", coins: []int{1, 0, 3}, amount: 6, wantErr: true},
		{name: "negative coin even for zero amount", coins: []int{1, -3}, amount: 0, wantErr: true},
	}
	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					before := slices.Clone(tt.coins)
					got, err := implementation.find(tt.coins, tt.amount)
					if (err != nil) != tt.wantErr {
						t.Fatalf("find(%v, %d) error = %v; want error %t", tt.coins, tt.amount, err, tt.wantErr)
					}
					if got != tt.want {
						t.Errorf("find(%v, %d) = %d; want %d", tt.coins, tt.amount, got, tt.want)
					}
					if !slices.Equal(tt.coins, before) {
						t.Errorf("find changed coins: got %v; want %v", tt.coins, before)
					}
				})
			}
		})
	}
}
