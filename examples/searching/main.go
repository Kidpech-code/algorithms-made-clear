package main

import (
	"fmt"

	"github.com/Kidpech-code/algorithms-made-clear/searching"
)

func main() {
	unsorted := []int{5, 1, 4, 2, 8, 3}
	sorted := []int{1, 2, 3, 4, 5, 8}
	target := 3

	fmt.Println("Unsorted:", unsorted)
	fmt.Println("Sorted:  ", sorted)
	fmt.Printf("Linear search for %d: index %d\n", target, searching.LinearSearch(unsorted, target))
	fmt.Printf("Binary search for %d: index %d\n", target, searching.BinarySearch(sorted, target))
	fmt.Printf("Missing value 6: linear %d, binary %d\n",
		searching.LinearSearch(unsorted, 6), searching.BinarySearch(sorted, 6))
}
