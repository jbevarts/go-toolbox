package dfs

import (
	ds "github.com/jbevarts/go-toolbox/datastructures"
)

func RecursiveDFS(root *ds.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{root.Val}
	result = append(result, RecursiveDFS(root.Left)...)
	result = append(result, RecursiveDFS(root.Right)...)

	return result
}
