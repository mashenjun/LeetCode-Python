package leafSimilarTrees

import (
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 思路：利用dfs找到所有的叶子节点
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaves1 := Leaves(root1)
	leaves2 := Leaves(root2)
	if len(leaves1) != len(leaves2) {
		return false
	}
	for i, v := range leaves1 {
		if v != leaves2[i] {
			return false
		}
	}
	return true
}

func Leaves(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ret := make([]int, 0)
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if curr.Right == nil && curr.Left == nil {
			ret = append(ret, curr.Val)
		}
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
	}
	return ret
}
