package leetcode

// A tree is an undirected graph in which any two vertices are connected by exactly one path. In other words, any connected graph without simple cycles is a tree.

// Given a tree of n nodes labelled from 0 to n - 1, and an array of n - 1 edges where edges[i] = [ai, bi] indicates that there is an undirected edge between the two nodes ai and bi in the tree, you can choose any node of the tree as the root. When you select a node x as the root, the result tree has height h. Among all possible rooted trees, those with minimum height (i.e. min(h))  are called minimum height trees (MHTs).

// Return a list of all MHTs' root labels. You can return the answer in any order.

// The height of a rooted tree is the number of edges on the longest downward path between the root and a leaf.
import (
	"container/list"
	"fmt"
	"slices"
)

// This implementation was too slow.  There is an algorithm for trimming leaves repeatedly to find the center nodes,
// which all graphs have exactly 2 of.
func findMinHeightTrees(n int, edges [][]int) []int {

	if len(edges) == 0 {
		return []int{0}
	}

	// for each node, run dfs iteratively, keeping track of the largest value on the stack.
	// return all values that equal the min height.

	// first, make a graph for the edge list

	graph := map[int][]int{}
	for _, e := range edges {
		if _, ok := graph[e[0]]; ok {
			graph[e[0]] = append(graph[e[0]], e[1])
		} else {
			graph[e[0]] = []int{e[1]}
		}

		if _, ok := graph[e[1]]; ok {
			graph[e[1]] = append(graph[e[1]], e[0])
		} else {
			graph[e[1]] = []int{e[0]}
		}
	}

	min := 0
	minNodes := []int{}
	for node, _ := range graph {
		height := bfsIterativeHeight(node, graph, min)
		fmt.Printf("height %v Node %v\n", height, node)
		if height == -1 {
			continue
		}
		if min == 0 {
			min = height
			minNodes = append(minNodes, node)
		} else {
			if height == min {
				minNodes = append(minNodes, node)
			} else if height < min {
				min = height
				minNodes = []int{node}
			}
		}
	}

	return minNodes
}

func bfsIterativeHeight(node int, graph map[int][]int, minSoFar int) int {
	q := list.New()
	q.PushBack(node)
	visited := []int{node}
	totalRows := 0

	for {
		if q.Len() == 0 {
			return totalRows
		}

		nextRow := list.New()
		l := q.Len() - 1
		for i := 0; i <= l; i++ {
			next := q.Front()
			v := next.Value.(int)
			q.Remove(next)

			for _, c := range graph[v] {
				if !slices.Contains(visited, c) {
					nextRow.PushBack(c)
					visited = append(visited, c)
				}
			}
		}

		if nextRow.Len() > 0 {
			totalRows += 1
		}

		if minSoFar != 0 && totalRows > minSoFar {
			return -1
		}

		q = nextRow
	}

	return totalRows

}
