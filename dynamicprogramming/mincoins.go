// Package dynamicprogramming shows how to reuse answers to smaller problems.
package dynamicprogramming

import "fmt"

// MinCoinsMemo returns the fewest coins needed to make amount using unlimited
// copies of each denomination. It returns -1 when amount is impossible.
// Negative amounts and nonpositive denominations are invalid.
// It takes O(amount * len(coins)) time and O(amount) memo and call-stack space.
func MinCoinsMemo(coins []int, amount int) (int, error) {
	if err := validate(coins, amount); err != nil {
		return 0, err
	}

	memo := make([]int, amount+1) // Zero means "not calculated" for positive amounts.
	var solve func(int) int
	solve = func(remaining int) int {
		if remaining == 0 {
			return 0
		}
		if memo[remaining] != 0 {
			return memo[remaining]
		}
		best := -1
		for _, coin := range coins {
			if coin > remaining {
				continue
			}
			smaller := solve(remaining - coin)
			if smaller >= 0 && (best < 0 || smaller+1 < best) {
				best = smaller + 1
			}
		}
		memo[remaining] = best
		return best
	}
	return solve(amount), nil
}

// MinCoinsTable solves the same problem by filling answers from zero through
// amount. It takes O(amount * len(coins)) time and O(amount) space.
func MinCoinsTable(coins []int, amount int) (int, error) {
	if err := validate(coins, amount); err != nil {
		return 0, err
	}

	dp := make([]int, amount+1) // dp[0] = 0: zero coins make zero.
	for target := 1; target <= amount; target++ {
		best := -1
		for _, coin := range coins {
			if coin > target || dp[target-coin] < 0 {
				continue
			}
			candidate := dp[target-coin] + 1
			if best < 0 || candidate < best {
				best = candidate
			}
		}
		dp[target] = best
	}
	return dp[amount], nil
}

func validate(coins []int, amount int) error {
	if amount < 0 {
		return fmt.Errorf("amount must be nonnegative: %d", amount)
	}
	for _, coin := range coins {
		if coin <= 0 {
			return fmt.Errorf("coin value must be positive: %d", coin)
		}
	}
	return nil
}
