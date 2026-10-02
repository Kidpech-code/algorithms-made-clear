package main

import (
	"fmt"

	"github.com/Kidpech-code/algorithms-made-clear/dynamicprogramming"
)

func main() {
	coins := []int{1, 3, 4}
	amount := 6
	memo, err := dynamicprogramming.MinCoinsMemo(coins, amount)
	if err != nil {
		panic(err)
	}
	table, err := dynamicprogramming.MinCoinsTable(coins, amount)
	if err != nil {
		panic(err)
	}
	fmt.Println("Token values:", coins)
	fmt.Println("Target:", amount)
	fmt.Printf("Memoization: %d tokens\n", memo)
	fmt.Printf("Tabulation: %d tokens\n", table)
}
