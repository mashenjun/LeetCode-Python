package symmetricTree

import (
	"testing"
)

// 101. Symmetric Tree

// Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).
// For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归检查是否对称
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricRecursive(root.Left, root.Right)
}


func isSymmetricRecursive(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}else if l == nil {
		return false
	}else if r == nil {
		return false
	}else if l.Val != r.Val {
		return false
	}
	return isSymmetricRecursive(l.Left, r.Right) && isSymmetricRecursive(l.Right, r.Left)
}

func Test_isSymmetric(t *testing.T) {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   3,
			},
			Right: &TreeNode{
				Val:   4,
			},
		},
		Right:&TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   4,
			},
			Right: &TreeNode{
				Val:   3,
			},
		},
	}
	isSymmetricRecursive(root.Left, root.Right)
}


