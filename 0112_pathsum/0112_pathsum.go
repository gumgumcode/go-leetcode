package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	// Check if it's a leaf node and if the sum matches the leaf node's value.
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}

	// Recursively check the left and right subtrees with the reduced sum.
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}

func main() {
	// Create a binary tree for testing.
	//     5
	//    / \
	//   4   8
	//  /   / \
	// 11  13  4
	// /  \      \
	// 7   2      1
	root := &TreeNode{Val: 5,
		Left: &TreeNode{Val: 4,
			Left: &TreeNode{Val: 11,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 2},
			},
		},
		Right: &TreeNode{Val: 8,
			Left: &TreeNode{Val: 13},
			Right: &TreeNode{Val: 4,
				Right: &TreeNode{Val: 1},
			},
		},
	}

	sum := 22
	result := hasPathSum(root, sum)
	fmt.Printf("Has path with sum %d: %v\n", sum, result)
}
