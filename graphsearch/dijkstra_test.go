package graphsearch_test

import (
	"math"
	"reflect"
	"slices"
	"strings"
	"testing"

	"github.com/Kidpech-code/algorithms-made-clear/graphsearch"
)

func weightedSampleGraph() map[int][]graphsearch.WeightedEdge {
	return map[int][]graphsearch.WeightedEdge{
		5: {{To: 1, Cost: 2}, {To: 4, Cost: 6}},
		1: {{To: 5, Cost: 2}, {To: 2, Cost: 2}, {To: 8, Cost: 3}},
		4: {{To: 5, Cost: 6}, {To: 3, Cost: 5}},
		2: {{To: 1, Cost: 2}},
		8: {{To: 1, Cost: 3}, {To: 3, Cost: 2}},
		3: {{To: 4, Cost: 5}, {To: 8, Cost: 2}},
	}
}

func TestDijkstra(t *testing.T) {
	tests := []struct {
		name     string
		graph    map[int][]graphsearch.WeightedEdge
		start    int
		goal     int
		wantCost int
		wantPath []int
	}{
		{name: "sample chooses faster three-edge route", graph: weightedSampleGraph(), start: 5, goal: 3, wantCost: 7, wantPath: []int{5, 1, 8, 3}},
		{name: "later improvement replaces expensive route", graph: map[int][]graphsearch.WeightedEdge{1: {{To: 2, Cost: 10}, {To: 3, Cost: 1}}, 3: {{To: 2, Cost: 1}, {To: 4, Cost: 100}}, 2: {{To: 4, Cost: 1}}}, start: 1, goal: 4, wantCost: 3, wantPath: []int{1, 3, 2, 4}},
		{name: "overflowing detour does not hide finite goal", graph: map[int][]graphsearch.WeightedEdge{1: {{To: 3, Cost: 3}, {To: 2, Cost: 1}}, 2: {{To: 4, Cost: math.MaxInt}}}, start: 1, goal: 3, wantCost: 3, wantPath: []int{1, 3}},
		{name: "unreachable despite overflowing detour", graph: map[int][]graphsearch.WeightedEdge{1: {{To: 2, Cost: 1}}, 2: {{To: 4, Cost: math.MaxInt}}}, start: 1, goal: 3, wantCost: -1},
		{name: "unreachable", graph: map[int][]graphsearch.WeightedEdge{1: {{To: 2, Cost: 4}}, 3: nil}, start: 1, goal: 3, wantCost: -1},
		{name: "isolated start", start: 1, goal: 3, wantCost: -1},
		{name: "start is goal", graph: weightedSampleGraph(), start: 5, goal: 5, wantCost: 0, wantPath: []int{5}},
		{name: "zero-cost cycle", graph: map[int][]graphsearch.WeightedEdge{1: {{To: 2, Cost: 0}}, 2: {{To: 1, Cost: 0}, {To: 3, Cost: 1}}}, start: 1, goal: 3, wantCost: 1, wantPath: []int{1, 2, 3}},
		{name: "directed edge", graph: map[int][]graphsearch.WeightedEdge{1: {{To: 2, Cost: 3}}}, start: 2, goal: 1, wantCost: -1},
		{name: "maximum representable cost", graph: map[int][]graphsearch.WeightedEdge{1: {{To: 2, Cost: math.MaxInt}}}, start: 1, goal: 2, wantCost: math.MaxInt, wantPath: []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			before := make(map[int][]graphsearch.WeightedEdge, len(tt.graph))
			for vertex, edges := range tt.graph {
				before[vertex] = slices.Clone(edges)
			}
			cost, path, err := graphsearch.Dijkstra(tt.graph, tt.start, tt.goal)
			if err != nil {
				t.Fatal(err)
			}
			if cost != tt.wantCost || !slices.Equal(path, tt.wantPath) || (path == nil) != (tt.wantPath == nil) {
				t.Fatalf("cost = %d, path = %v; want cost = %d, path = %v", cost, path, tt.wantCost, tt.wantPath)
			}
			if tt.graph != nil && !reflect.DeepEqual(tt.graph, before) {
				t.Fatal("Dijkstra changed its input graph")
			}
		})
	}
}

func TestDijkstraRejectsInvalidWeights(t *testing.T) {
	tests := []struct {
		name    string
		graph   map[int][]graphsearch.WeightedEdge
		start   int
		goal    int
		message string
	}{
		{name: "negative edge in disconnected component", graph: map[int][]graphsearch.WeightedEdge{1: {{To: 2, Cost: 1}}, 9: {{To: 10, Cost: -1}}}, start: 1, goal: 2, message: "negative"},
		{name: "negative edge when start is goal", graph: map[int][]graphsearch.WeightedEdge{1: {{To: 2, Cost: -1}}}, start: 1, goal: 1, message: "negative"},
		{name: "cost overflow", graph: map[int][]graphsearch.WeightedEdge{1: {{To: 2, Cost: math.MaxInt}}, 2: {{To: 3, Cost: 1}}}, start: 1, goal: 3, message: "overflow"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := graphsearch.Dijkstra(tt.graph, tt.start, tt.goal)
			if err == nil || !strings.Contains(err.Error(), tt.message) {
				t.Fatalf("error = %v; want message containing %q", err, tt.message)
			}
		})
	}
}
