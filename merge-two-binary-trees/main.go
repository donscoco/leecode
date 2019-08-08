package main

import "fmt"

//https://leetcode.com/problems/merge-two-binary-trees

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	return mergeNode(t1,t2)
}

func mergeNode(n1 *TreeNode, n2 *TreeNode) *TreeNode {
	if n1 == nil && n2 == nil {
		return nil
	}
	if n1 != nil && n2 == nil {
		return n1
	}
	if n1 == nil && n2 != nil {
		return n2
	}
	n1.Left = mergeNode(n1.Left,n2.Left)
	n1.Right = mergeNode(n1.Right,n2.Right)
	n1.Val = n1.Val + n2.Val
	return n1
}
func main() {
	var n *TreeNode
	if n == nil {
		fmt.Println("hw")
	}
}
