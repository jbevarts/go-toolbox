package dfs

import (
	"fmt"
	"testing"

	ds "github.com/jbevarts/go-toolbox/datastructures"
	examples "github.com/jbevarts/go-toolbox/examples"
	"github.com/stretchr/testify/assert"
)

func Test_iterativeDFS(t *testing.T) {
	testCases := []struct {
		root       *ds.TreeNode
		visitOrder []int
	}{
		{
			root:       examples.Tree1,
			visitOrder: []int{1, 2, 4, 8, 9, 5, 10, 11, 3, 6, 12, 13, 7, 14, 15},
		},
		{
			root:       examples.Tree2,
			visitOrder: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		},
		{
			root:       examples.Tree3,
			visitOrder: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		},
		{
			root:       examples.Tree4,
			visitOrder: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		},
		{
			root:       examples.Tree5,
			visitOrder: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 115, 112, 109, 106, 103},
		},
	}
	for _, tt := range testCases {
		order := IterativeDFS(tt.root)
		fmt.Println(order)

		assert.Equal(t, tt.visitOrder, order, "Order incorrect")
	}
}

func Test_iterativeDFSInorder(t *testing.T) {
	testCases := []struct {
		root       *ds.TreeNode
		visitOrder []int
	}{
		{
			root:       examples.Tree1,
			visitOrder: []int{8, 4, 9, 2, 10, 5, 11, 1, 12, 6, 13, 3, 14, 7, 15},
		},
		{
			root:       examples.Tree2,
			visitOrder: []int{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			root:       examples.Tree3,
			visitOrder: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		},
		{
			root:       examples.Tree4,
			visitOrder: []int{2, 4, 6, 8, 10, 12, 14, 15, 13, 11, 9, 7, 5, 3, 1},
		},
		{
			root:       examples.Tree5,
			visitOrder: []int{15, 115, 14, 13, 12, 112, 11, 10, 9, 109, 8, 7, 6, 106, 5, 4, 3, 103, 2, 1},
		},
	}
	for _, tt := range testCases {
		order := IterativeDFSInorder(tt.root)
		fmt.Println(order)

		assert.Equal(t, tt.visitOrder, order, "Order incorrect")
	}
}
