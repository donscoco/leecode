package range_sum_of_bst


//https://leetcode.com/problems/range-sum-of-bst/


type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	result := sum(L,R,root)
	return result
}

func sum(L,R int,node *TreeNode) int{
	if node == nil {
		return 0
	}
	if node.Val<L{
		return sum(L,R,node.Right)
	}
	if node.Val>R{
		return sum(L,R,node.Left)
	}
	return node.Val + sum(L,R,node.Left) + sum(L,R,node.Right)
}

func main() {
	
}
