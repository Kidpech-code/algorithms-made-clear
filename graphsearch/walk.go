// Package graphsearch shows how breadth-first and depth-first search walk a graph.
package graphsearch

import "slices"

// BFS visits every vertex reachable from start in breadth-first order.
// It also returns a shortest path to goal by number of edges, or nil if unreachable.
// Neighbor slice order determines the order of visits and which shortest path is returned.
func BFS(graph map[int][]int, start, goal int) (order, path []int) {
	queue := []int{start}
	parent := map[int]int{start: start}
	for head := 0; head < len(queue); head++ {
		vertex := queue[head]
		order = append(order, vertex)
		for _, neighbor := range graph[vertex] {
			if _, seen := parent[neighbor]; seen {
				continue
			}
			parent[neighbor] = vertex
			queue = append(queue, neighbor)
		}
	}
	return order, pathTo(parent, start, goal)
}

// DFS visits every vertex reachable from start in depth-first preorder.
// Its path to goal follows first-discovery edges and need not be shortest.
// Neighbor slice order determines the order of visits.
func DFS(graph map[int][]int, start, goal int) (order, path []int) {
	parent := map[int]int{start: start}
	var visit func(int)
	visit = func(vertex int) {
		order = append(order, vertex)
		for _, neighbor := range graph[vertex] {
			if _, seen := parent[neighbor]; seen {
				continue
			}
			parent[neighbor] = vertex
			visit(neighbor)
		}
	}
	visit(start)
	return order, pathTo(parent, start, goal)
}

func pathTo(parent map[int]int, start, goal int) []int {
	if _, found := parent[goal]; !found {
		return nil
	}
	var path []int
	for vertex := goal; ; vertex = parent[vertex] {
		path = append(path, vertex)
		if vertex == start {
			break
		}
	}
	slices.Reverse(path)
	return path
}
