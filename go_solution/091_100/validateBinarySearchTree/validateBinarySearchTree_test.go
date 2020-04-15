package validateBinarySearchTree

import "testing"

// 098. Validate Binary Search Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return validate(root.Right, &root.Val, nil) && validate(root.Left, nil, &root.Val)
}

func validate(root *TreeNode, lower *int, upper *int) bool {
	if root == nil {
		return true
	}
	if lower != nil && root.Val < *lower {
		return false
	}
	if upper != nil && root.Val > *upper {
		return false
	}
	return validate(root.Right, &root.Val, upper) && validate(root.Left, lower, &root.Val)
}

func TestIsValidBST(t *testing.T) {
	root := &TreeNode{
		Val:   2,
		Left:  &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	t.Log(isValidBST(root))
} 
