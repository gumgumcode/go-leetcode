package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	// Find the middle element of the sorted array
	mid := len(nums) / 2

	// Create a new tree node with the middle element
	root := &TreeNode{Val: nums[mid]}

	// Recursively construct the left and right subtrees
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])

	return root
}

// Helper function to print the binary tree (in-order traversal)
func printTree(root *TreeNode) {
	if root != nil {
		printTree(root.Left)
		fmt.Printf("%d ", root.Val)
		printTree(root.Right)
	}
}

func main() {
	// Example input: a sorted array
	nums := []int{-10, -3, 0, 5, 9}

	// Convert the sorted array to a binary search tree
	result := sortedArrayToBST(nums)

	// Print the binary tree (in-order traversal)
	fmt.Print("In-order traversal of the binary search tree: ")
	printTree(result)
	fmt.Println()
}
