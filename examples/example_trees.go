package examples

import (
	ds "github.com/jbevarts/go-toolbox/datastructures"
)

/*
Tree1: Complete Binary Tree (15 nodes)

	           1
	       ┌───┴───┐
	    2           3
	 ┌──┴──┐     ┌──┴──┐
	 4     5     6     7
	┌┴┐   ┌┴┐   ┌┴┐   ┌┴┐
	8 9  10 11 12 13 14 15
*/
var Tree1 = &ds.TreeNode{Val: 1,
	Left: &ds.TreeNode{Val: 2,
		Left: &ds.TreeNode{Val: 4,
			Left:  &ds.TreeNode{Val: 8},
			Right: &ds.TreeNode{Val: 9},
		},
		Right: &ds.TreeNode{Val: 5,
			Left:  &ds.TreeNode{Val: 10},
			Right: &ds.TreeNode{Val: 11},
		},
	},
	Right: &ds.TreeNode{Val: 3,
		Left: &ds.TreeNode{Val: 6,
			Left:  &ds.TreeNode{Val: 12},
			Right: &ds.TreeNode{Val: 13},
		},
		Right: &ds.TreeNode{Val: 7,
			Left:  &ds.TreeNode{Val: 14},
			Right: &ds.TreeNode{Val: 15},
		},
	},
}

/*
Tree2: Left-Skewed Tree (15 nodes) - Each node only has LEFT children

	                            1
	                           /
	                          2
	                         /
	                        3
	                       /
	                      4
	                     /
	                    5
	                   /
	                  6
	                 /
	                7
	               /
	              8
	             /
	            9
	           /
	          10
	         /
	        11
	       /
	      12
	     /
	    13
	   /
	  14
	 /
	15
*/
var Tree2 = func() *ds.TreeNode {
	t := &ds.TreeNode{Val: 1}
	cur := t
	for i := 2; i <= 15; i++ {
		cur.Left = &ds.TreeNode{Val: i}
		cur = cur.Left
	}
	return t
}()

/*
Tree3: Right-Skewed Tree (15 nodes) - Each node only has RIGHT children

1

	\
	 2
	  \
	   3
	    \
	     4
	      \
	       5
	        \
	         6
	          \
	           7
	            \
	             8
	              \
	               9
	                \
	                 10
	                  \
	                   11
	                    \
	                     12
	                      \
	                       13
	                        \
	                         14
	                          \
	                           15
*/
var Tree3 = func() *ds.TreeNode {
	t := &ds.TreeNode{Val: 1}
	cur := t
	for i := 2; i <= 15; i++ {
		cur.Right = &ds.TreeNode{Val: i}
		cur = cur.Right
	}
	return t
}()

/*
Tree4: Alternating Left-Right Zigzag Chain (15 nodes)

	  1
	 /
	2
	 \
	  3
	 /
	4
	 \
	  5
	 /
	6
	 \
	  7
	 /
	8
	 \
	  9
	 /
	10
	 \
	  11
	 /
	12
	 \
	  13
	 /
	14
	 \
	  15
*/
var Tree4 = func() *ds.TreeNode {
	t := &ds.TreeNode{Val: 1}
	cur := t
	for i := 2; i <= 15; i++ {
		if i%2 == 0 {
			cur.Left = &ds.TreeNode{Val: i}
			cur = cur.Left
		} else {
			cur.Right = &ds.TreeNode{Val: i}
			cur = cur.Right
		}
	}
	return t
}()

/*
Tree5: Left-Heavy Tree with Right Branches (15 main nodes + 5 branch nodes)

	                           1
	                          /
	                         2
	                        /
	                       3
	                      / \
	                     4   103
	                    /
	                   5
	                  /
	                 6
	                / \
	               7   106
	              /
	             8
	            /
	           9
	          / \
	         10  109
	        /
	       11
	      /
	     12
	    / \
	   13  112
	  /
	 14
	/

15

	\
	 115
*/
var Tree5 = func() *ds.TreeNode {
	t := &ds.TreeNode{Val: 1}
	cur := t
	for i := 2; i <= 15; i++ {
		cur.Left = &ds.TreeNode{Val: i}
		if i%3 == 0 {
			cur.Left.Right = &ds.TreeNode{Val: i + 100}
		}
		cur = cur.Left
	}
	return t
}()

/*
Bst1: Valid BST - Balanced Tree (7 nodes)

	     8
	   /   \
	  3     10
	 / \      \
	1   6      14
	   / \    /
	  4   7  13
*/
var Bst1 = &ds.TreeNode{Val: 8,
	Left: &ds.TreeNode{Val: 3,
		Left: &ds.TreeNode{Val: 1},
		Right: &ds.TreeNode{Val: 6,
			Left:  &ds.TreeNode{Val: 4},
			Right: &ds.TreeNode{Val: 7},
		},
	},
	Right: &ds.TreeNode{Val: 10,
		Right: &ds.TreeNode{Val: 14,
			Left: &ds.TreeNode{Val: 13},
		},
	},
}

/*
Bst2: Valid BST - Left-Heavy Tree (8 nodes)

	        15
	       /  \
	      10   20
	     /  \
	    6    12
	   / \     \
	  3   8     13
	 /
	1
*/
var Bst2 = &ds.TreeNode{Val: 15,
	Left: &ds.TreeNode{Val: 10,
		Left: &ds.TreeNode{Val: 6,
			Left: &ds.TreeNode{Val: 3,
				Left: &ds.TreeNode{Val: 1},
			},
			Right: &ds.TreeNode{Val: 8},
		},
		Right: &ds.TreeNode{Val: 12,
			Right: &ds.TreeNode{Val: 13},
		},
	},
	Right: &ds.TreeNode{Val: 20},
}

/*
Bst3: Valid BST - Right-Heavy Tree (8 nodes)

	    5
	   / \
	  2   12
	 /   /  \
	1   8    18
	   / \   /  \
	  6  10 16  25
*/
var Bst3 = &ds.TreeNode{Val: 5,
	Left: &ds.TreeNode{Val: 2,
		Left: &ds.TreeNode{Val: 1},
	},
	Right: &ds.TreeNode{Val: 12,
		Left: &ds.TreeNode{Val: 8,
			Left:  &ds.TreeNode{Val: 6},
			Right: &ds.TreeNode{Val: 10},
		},
		Right: &ds.TreeNode{Val: 18,
			Left:  &ds.TreeNode{Val: 16},
			Right: &ds.TreeNode{Val: 25},
		},
	},
}

/*
Bst4: Valid BST - Linear Ascending Chain (8 nodes)

	50
	 \
	  60
	   \
	    70
	     \
	      80
	       \
	        90
	         \
	          100
	           \
	            110
	             \
	              120
*/
var Bst4 = func() *ds.TreeNode {
	values := []int{50, 60, 70, 80, 90, 100, 110, 120}
	root := &ds.TreeNode{Val: values[0]}
	cur := root
	for i := 1; i < len(values); i++ {
		cur.Right = &ds.TreeNode{Val: values[i]}
		cur = cur.Right
	}
	return root
}()

/*
Bst5: Valid BST - Linear Descending Chain (8 nodes)

	              50
	             /
	            40
	           /
	          30
	         /
	        20
	       /
	      10
	     /
	    5
	   /
	  2
	 /
	1
*/
var Bst5 = func() *ds.TreeNode {
	values := []int{50, 40, 30, 20, 10, 5, 2, 1}
	root := &ds.TreeNode{Val: values[0]}
	cur := root
	for i := 1; i < len(values); i++ {
		cur.Left = &ds.TreeNode{Val: values[i]}
		cur = cur.Left
	}
	return root
}()
