package main

import (
	"fmt"

	"github.com/Kidpech-code/algorithms-made-clear/recursion"
)

func main() {
	values := []int{5, 1, 4, 2, 8, 3}
	fmt.Println("Values:", values)
	fmt.Println("Recursive sum:", recursion.Sum(values))
	fmt.Println("Empty slice:", recursion.Sum(nil))
}
