package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// both does not exist
	if p == nil && q == nil {
		return true
	}
	// one of them exists, other one doesn't
	if p == nil || q == nil {
		return false
	}
	// both exists
	if p.Val != q.Val {
		return false
	}
	// check subroots
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func case1() {
	p := &TreeNode{Val: 1}
	p.Left = &TreeNode{Val: 2}
	p.Right = &TreeNode{Val: 3}

	q := &TreeNode{Val: 1}
	q.Left = &TreeNode{Val: 2}
	q.Right = &TreeNode{Val: 3}

	result := isSameTree(p, q)
	fmt.Println(result)
}

func case2() {
	p := &TreeNode{Val: 1}
	p.Left = &TreeNode{Val: 2}
	p.Right = &TreeNode{Val: 3}

	q := &TreeNode{Val: 1}
	q.Left = &TreeNode{Val: 2}
	q.Right = &TreeNode{Val: 4}

	result := isSameTree(p, q)
	fmt.Println(result)
}

func main() {
	case1()
	case2()
}
