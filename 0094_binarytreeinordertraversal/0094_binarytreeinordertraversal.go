package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// iterative!
func inorderTraversal(root *TreeNode) []int {
	var result []int
	stack := []*TreeNode{}

	curr := root

	for curr != nil || len(stack) > 0 {
		// traverse the left side of the tree
		// and add all the nodes along the way to our stack.
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, curr.Val)
		curr = curr.Right
	}

	return result
}

// recursive!
func inorderTraversal2(root *TreeNode) []int {
	var result []int
	helper(root, &result)
	return result
}

func helper(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	helper(root.Left, result)
	*result = append(*result, root.Val)
	helper(root.Right, result)
}

func main() {
	// Create a binary tree
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}

	// Perform inorder traversal
	result := inorderTraversal2(root)

	fmt.Println(result) // Print the result
}

// preorder: root left right
// inorder: left root right
// postorder: left right root

// 1
//  \
//   2
//  /
// 3

// inorder: 1 -> 3 -> 2
