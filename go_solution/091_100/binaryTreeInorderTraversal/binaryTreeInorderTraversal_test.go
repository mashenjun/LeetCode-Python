package binaryTreeInorderTraversal

import "testing"

func Test_MorrisTraversal(t *testing.T) {
	root := &TreeNode{
		Val:   1,
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   3,

			},
		},
	}
	ret := morrisTraversal(root)
	t.Log(ret)
}
