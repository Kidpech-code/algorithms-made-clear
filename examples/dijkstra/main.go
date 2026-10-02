package main

import (
	"fmt"
	"log"

	"github.com/Kidpech-code/algorithms-made-clear/graphsearch"
)

func main() {
	// Every corridor appears twice because walking is possible in both directions.
	graph := map[int][]graphsearch.WeightedEdge{
		5: {{To: 1, Cost: 2}, {To: 4, Cost: 6}},
		1: {{To: 5, Cost: 2}, {To: 2, Cost: 2}, {To: 8, Cost: 3}},
		4: {{To: 5, Cost: 6}, {To: 3, Cost: 5}},
		2: {{To: 1, Cost: 2}},
		8: {{To: 1, Cost: 3}, {To: 3, Cost: 2}},
		3: {{To: 4, Cost: 5}, {To: 8, Cost: 2}},
	}

	// BFS ignores travel times, so it chooses a route with the fewest doors.
	unweighted := make(map[int][]int, len(graph))
	for room, corridors := range graph {
		for _, corridor := range corridors {
			unweighted[room] = append(unweighted[room], corridor.To)
		}
	}
	_, bfsPath := graphsearch.BFS(unweighted, 5, 3)
	cost, path, err := graphsearch.Dijkstra(graph, 5, 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("BFS (fewest doors):", bfsPath)
	fmt.Println("Dijkstra (least time):", path)
	fmt.Printf("Travel time: %d minutes\n", cost)
}
