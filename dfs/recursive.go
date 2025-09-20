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

func RecursiveDFSInorder(root *ds.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := RecursiveDFSInorder(root.Left)
	result = append(result, root.Val)
	result = append(result, RecursiveDFSInorder(root.Right)...)

	return result
}

func RecursiveDFSPostorder(root *ds.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := RecursiveDFSPostorder(root.Left)
	result = append(result, RecursiveDFSPostorder(root.Right)...)
	result = append(result, root.Val)

	return result
}
