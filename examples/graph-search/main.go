package main

import (
	"fmt"

	"github.com/Kidpech-code/algorithms-made-clear/graphsearch"
)

func main() {
	graph := map[int][]int{
		5: {1, 4},
		1: {5, 2, 8},
		4: {5, 3},
		2: {1},
		8: {1, 3},
		3: {4, 8},
	}
	fmt.Println("Edges: 5-1, 5-4, 1-2, 1-8, 4-3, 8-3")
	bfsOrder, bfsPath := graphsearch.BFS(graph, 5, 3)
	fmt.Println("BFS order from 5:", bfsOrder)
	fmt.Println("BFS path to 3:", bfsPath)
	dfsOrder, dfsPath := graphsearch.DFS(graph, 5, 3)
	fmt.Println("DFS order from 5:", dfsOrder)
	fmt.Println("DFS path to 3:", dfsPath)
}
