package binaryTreePreorderTraversal

import "testing"

func Test_PreorderTraversal(t *testing.T) {
	root := &TreeNode{
		Val:   1,
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   3,
			},
		},
	}
	t.Log(preorderMorris(root))
}
