package dfs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_iterativeDFS(t *testing.T) {
	testCases := []struct {
		root       *TreeNode
		visitOrder []int
	}{
		{
			root:       tree1,
			visitOrder: []int{1, 2, 4, 8, 9, 5, 10, 11, 3, 6, 12, 13, 7, 14, 15},
		},
		{
			root:       tree2,
			visitOrder: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		},
		{
			root:       tree3,
			visitOrder: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		},
		{
			root:       tree4,
			visitOrder: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		},
		{
			root:       tree5,
			visitOrder: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 115, 112, 109, 106, 103},
		},
	}
	for _, tt := range testCases {
		order := IterativeDFS(tt.root)
		fmt.Println(order)

		assert.Equal(t, tt.visitOrder, order, "Order incorrect")
	}
}
