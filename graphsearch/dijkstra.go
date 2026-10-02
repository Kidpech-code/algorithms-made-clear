package graphsearch

import (
	"container/heap"
	"fmt"
	"math"
)

// WeightedEdge is a directed edge to another vertex with a nonnegative cost.
type WeightedEdge struct {
	To   int
	Cost int
}

// Dijkstra returns the least-cost route from start to goal.
// List both directions to represent an undirected edge. Negative costs are
// rejected, even in disconnected components. An unreachable goal returns -1
// and a nil path; a reachable goal whose least cost exceeds int returns an error.
func Dijkstra(graph map[int][]WeightedEdge, start, goal int) (cost int, path []int, err error) {
	for from, edges := range graph {
		for _, edge := range edges {
			if edge.Cost < 0 {
				return 0, nil, fmt.Errorf("negative cost on edge %d -> %d", from, edge.To)
			}
		}
	}

	best := map[int]int{start: 0}
	parent := map[int]int{start: start}
	queue := &routeQueue{{vertex: start, cost: 0}}
	heap.Init(queue)
	overflowed := false

	for queue.Len() > 0 {
		current := heap.Pop(queue).(routeStep)
		if current.cost != best[current.vertex] {
			continue // A better route reached this vertex after it entered the queue.
		}
		if current.vertex == goal {
			return current.cost, pathTo(parent, start, goal), nil
		}
		for _, edge := range graph[current.vertex] {
			if current.cost > math.MaxInt-edge.Cost {
				overflowed = true
				continue
			}
			next := current.cost + edge.Cost
			if previous, found := best[edge.To]; found && next >= previous {
				continue
			}
			best[edge.To] = next
			parent[edge.To] = current.vertex
			heap.Push(queue, routeStep{vertex: edge.To, cost: next})
		}
	}
	if overflowed {
		unweighted := make(map[int][]int, len(graph))
		for from, edges := range graph {
			for _, edge := range edges {
				unweighted[from] = append(unweighted[from], edge.To)
			}
		}
		if _, possible := BFS(unweighted, start, goal); possible != nil {
			return 0, nil, fmt.Errorf("shortest path cost overflows int")
		}
	}
	return -1, nil, nil
}

type routeStep struct {
	vertex int
	cost   int
}

type routeQueue []routeStep

func (q routeQueue) Len() int           { return len(q) }
func (q routeQueue) Less(i, j int) bool { return q[i].cost < q[j].cost }
func (q routeQueue) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }

func (q *routeQueue) Push(value any) { *q = append(*q, value.(routeStep)) }
func (q *routeQueue) Pop() any {
	last := len(*q) - 1
	value := (*q)[last]
	*q = (*q)[:last]
	return value
}
