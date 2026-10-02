package graphsearch_test

import (
	"reflect"
	"slices"
	"testing"

	"github.com/Kidpech-code/algorithms-made-clear/graphsearch"
)

func sampleGraph() map[int][]int {
	return map[int][]int{
		5: {1, 4},
		1: {5, 2, 8},
		4: {5, 3},
		2: {1},
		8: {1, 3},
		3: {4, 8},
	}
}

func TestWalks(t *testing.T) {
	tests := []struct {
		name      string
		walk      func(map[int][]int, int, int) ([]int, []int)
		graph     map[int][]int
		start     int
		goal      int
		wantOrder []int
		wantPath  []int
	}{
		{"BFS sample", graphsearch.BFS, sampleGraph(), 5, 3, []int{5, 1, 4, 2, 8, 3}, []int{5, 4, 3}},
		{"DFS sample", graphsearch.DFS, sampleGraph(), 5, 3, []int{5, 1, 2, 8, 3, 4}, []int{5, 1, 8, 3}},
		{"BFS disconnected", graphsearch.BFS, map[int][]int{5: {1}, 1: {5}, 3: {4}, 4: {3}}, 5, 3, []int{5, 1}, nil},
		{"DFS disconnected", graphsearch.DFS, map[int][]int{5: {1}, 1: {5}, 3: {4}, 4: {3}}, 5, 3, []int{5, 1}, nil},
		{"BFS start is goal", graphsearch.BFS, sampleGraph(), 5, 5, []int{5, 1, 4, 2, 8, 3}, []int{5}},
		{"DFS start is goal", graphsearch.DFS, sampleGraph(), 5, 5, []int{5, 1, 2, 8, 3, 4}, []int{5}},
		{"BFS isolated start", graphsearch.BFS, nil, 5, 3, []int{5}, nil},
		{"DFS isolated start", graphsearch.DFS, nil, 5, 3, []int{5}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var before map[int][]int
			if tt.graph != nil {
				before = make(map[int][]int, len(tt.graph))
				for vertex, neighbors := range tt.graph {
					before[vertex] = slices.Clone(neighbors)
				}
			}
			order, path := tt.walk(tt.graph, tt.start, tt.goal)
			if !slices.Equal(order, tt.wantOrder) || !slices.Equal(path, tt.wantPath) || (path == nil) != (tt.wantPath == nil) {
				t.Fatalf("order = %v, path = %v; want order = %v, path = %v", order, path, tt.wantOrder, tt.wantPath)
			}
			if !reflect.DeepEqual(tt.graph, before) {
				t.Fatal("walk changed its input graph")
			}
		})
	}
}
