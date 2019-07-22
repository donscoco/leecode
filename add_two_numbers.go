package main

import "fmt"

//https://leetcode.com/problems/add-two-numbers/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val int
	Next *ListNode
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var currentNode = &ListNode{}
	resultPre := currentNode
	var carry int
	for l1 != nil || l2 != nil{
		// plus option             tip:nil
		var v1,v2 int
		if l1 != nil {
			v1=l1.Val
		}
		if l2 != nil {
			v2=l2.Val
		}
		sum := carry + v1 + v2
		currentValue := sum%10
		carry = sum/10

		//current move
		currentNode.Next=&ListNode{
			Val:currentValue,
		}
		currentNode = currentNode.Next

		// next prepare      tip:nil
		if l1!=nil{
			l1=l1.Next
		}
		if l2!=nil{
			l2=l2.Next
		}
	}

	if carry > 0{
		currentNode.Next=&ListNode{
			Val:carry,
		}
	}
	return resultPre.Next
}

func main() {
	var l1,l2,cur *ListNode
	var arr1,arr2 = []int{2,4,3},[]int{5,6,4}
	l1=&ListNode{}
	l2=&ListNode{}
	cur=l1
	for _,v := range arr1 {
		cur.Next=&ListNode{}
		cur.Next.Val=v
		cur=cur.Next
	}
	cur=l2
	for _,v := range arr2 {
		cur.Next=&ListNode{}
		cur.Next.Val=v
		cur=cur.Next
	}
	result:=addTwoNumbers(l1.Next,l2.Next)
	for result!=nil{
		fmt.Println(result.Val)
		result=result.Next
	}
}
