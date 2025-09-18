package dfs

/*
tree1: Complete Binary Tree (15 nodes)

              1
          ┌───┴───┐
       2           3
    ┌──┴──┐     ┌──┴──┐
    4     5     6     7
   ┌┴┐   ┌┴┐   ┌┴┐   ┌┴┐
   8 9  10 11 12 13 14 15
*/
var tree1 = &TreeNode{Val: 1,
	Left: &TreeNode{Val: 2,
		Left: &TreeNode{Val: 4,
			Left:  &TreeNode{Val: 8},
			Right: &TreeNode{Val: 9},
		},
		Right: &TreeNode{Val: 5,
			Left:  &TreeNode{Val: 10},
			Right: &TreeNode{Val: 11},
		},
	},
	Right: &TreeNode{Val: 3,
		Left: &TreeNode{Val: 6,
			Left:  &TreeNode{Val: 12},
			Right: &TreeNode{Val: 13},
		},
		Right: &TreeNode{Val: 7,
			Left:  &TreeNode{Val: 14},
			Right: &TreeNode{Val: 15},
		},
	},
}

/*
tree2: Left-Skewed Tree (15 nodes) - Each node only has LEFT children

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
var tree2 = func() *TreeNode {
	t := &TreeNode{Val: 1}
	cur := t
	for i := 2; i <= 15; i++ {
		cur.Left = &TreeNode{Val: i}
		cur = cur.Left
	}
	return t
}()

/*
tree3: Right-Skewed Tree (15 nodes) - Each node only has RIGHT children

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
var tree3 = func() *TreeNode {
	t := &TreeNode{Val: 1}
	cur := t
	for i := 2; i <= 15; i++ {
		cur.Right = &TreeNode{Val: i}
		cur = cur.Right
	}
	return t
}()

/*
tree4: Alternating Left-Right Zigzag Chain (15 nodes)

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
var tree4 = func() *TreeNode {
	t := &TreeNode{Val: 1}
	cur := t
	for i := 2; i <= 15; i++ {
		if i%2 == 0 {
			cur.Left = &TreeNode{Val: i}
			cur = cur.Left
		} else {
			cur.Right = &TreeNode{Val: i}
			cur = cur.Right
		}
	}
	return t
}()

/*
tree5: Left-Heavy Tree with Right Branches (15 main nodes + 5 branch nodes)

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
var tree5 = func() *TreeNode {
	t := &TreeNode{Val: 1}
	cur := t
	for i := 2; i <= 15; i++ {
		cur.Left = &TreeNode{Val: i}
		if i%3 == 0 {
			cur.Left.Right = &TreeNode{Val: i + 100}
		}
		cur = cur.Left
	}
	return t
}()

/*
bst1: Valid BST - Balanced Tree (7 nodes)

       8
     /   \
    3     10
   / \      \
  1   6      14
     / \    /
    4   7  13
*/
var bst1 = &TreeNode{Val: 8,
	Left: &TreeNode{Val: 3,
		Left: &TreeNode{Val: 1},
		Right: &TreeNode{Val: 6,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 7},
		},
	},
	Right: &TreeNode{Val: 10,
		Right: &TreeNode{Val: 14,
			Left: &TreeNode{Val: 13},
		},
	},
}

/*
bst2: Valid BST - Left-Heavy Tree (8 nodes)

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
var bst2 = &TreeNode{Val: 15,
	Left: &TreeNode{Val: 10,
		Left: &TreeNode{Val: 6,
			Left: &TreeNode{Val: 3,
				Left: &TreeNode{Val: 1},
			},
			Right: &TreeNode{Val: 8},
		},
		Right: &TreeNode{Val: 12,
			Right: &TreeNode{Val: 13},
		},
	},
	Right: &TreeNode{Val: 20},
}

/*
bst3: Valid BST - Right-Heavy Tree (8 nodes)

     5
    / \
   2   12
  /   /  \
 1   8    18
    / \   /  \
   6  10 16  25
*/
var bst3 = &TreeNode{Val: 5,
	Left: &TreeNode{Val: 2,
		Left: &TreeNode{Val: 1},
	},
	Right: &TreeNode{Val: 12,
		Left: &TreeNode{Val: 8,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 10},
		},
		Right: &TreeNode{Val: 18,
			Left:  &TreeNode{Val: 16},
			Right: &TreeNode{Val: 25},
		},
	},
}

/*
bst4: Valid BST - Linear Ascending Chain (8 nodes)

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
var bst4 = func() *TreeNode {
	values := []int{50, 60, 70, 80, 90, 100, 110, 120}
	root := &TreeNode{Val: values[0]}
	cur := root
	for i := 1; i < len(values); i++ {
		cur.Right = &TreeNode{Val: values[i]}
		cur = cur.Right
	}
	return root
}()

/*
bst5: Valid BST - Linear Descending Chain (8 nodes)

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
var bst5 = func() *TreeNode {
	values := []int{50, 40, 30, 20, 10, 5, 2, 1}
	root := &TreeNode{Val: values[0]}
	cur := root
	for i := 1; i < len(values); i++ {
		cur.Left = &TreeNode{Val: values[i]}
		cur = cur.Left
	}
	return root
}()
