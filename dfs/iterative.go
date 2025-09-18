package dfs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IterativeDFS(root *TreeNode) []int {
	visited := []int{}
	stack := []*TreeNode{}

	cur := root

	for cur != nil || len(stack) > 0 {
		for cur != nil {
			visited = append(visited, cur.Val)
			stack = append(stack, cur)
			cur = cur.Left
		}

		next := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur = next.Right
	}
	return visited
}
