package kthSmallestElementInABST

// 230. Kth Smallest Element in a BST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 	思路: 用中序遍历k次
func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return -1
	}
	stack := []*TreeNode{}
	curr := root
	for len(stack) != 0 || curr != nil {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr, stack = stack[len(stack)-1], stack[:len(stack)-1] // pop form stack
		k--
		if k == 0 {
			return curr.Val
		}
		curr = curr.Right
	}
	return -1
}
