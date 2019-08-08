package main

import "fmt"

//https://leetcode.com/problems/di-string-match/
type Node struct {
	Val int
	Status byte
	Prev *Node
	Next *Node
}
func diStringMatch(S string) []int {

	result := make([]int,len(S)+1)
	op := []byte(S)
	node := new(Node)
	node.Val = 0
	node.Status = op[0]
	op = op[1:]
	header := node
	current := node
	for num := 1; num < len(S)+1; num++ {
		node := new(Node)
		node.Val = num
		if len(op) > 0 {
			node.Status = op[0]
			op = op[1:]
		}

		if current.Status == 73 { // increase
			current.Next,node.Prev = node,current
			current = current.Next
			continue
		}
		// decrease
		point := current
		for point.Prev != nil {
			point = point.Prev
			if point.Status == 73 {
				point.Prev.Next,node.Prev = node,point.Prev
				point.Prev,node.Next = node,point
				continue
			}
		}
		point.Prev,node.Next = node,point

	}
	for index := 0; header.Next != nil; index++ {
		result[index] = header.Val
		header = header.Next
	}
	return result
}
func main() {
	s := "ID"
	fmt.Println(diStringMatch(s))
}
