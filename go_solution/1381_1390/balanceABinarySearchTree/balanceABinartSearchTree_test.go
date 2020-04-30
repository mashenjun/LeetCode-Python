package balanceABinarySearchTree

import "testing"

// 1382. Balance a Binary Search Tree

// Given a binary search tree, return a balanced binary search tree with the same node values.
// A binary search tree is balanced if and only if the depth of the two subtrees of every node never differ by more than 1.
// If there is more than one answer, return any of them.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 思路：
// 1. 先中序遍历获取有序的数组
// 2. 再通过有序数组重新构造BST

func balanceBST(root *TreeNode) *TreeNode {
	vals := inOrderTraversal(root)
	return buildTree(vals, 0, len(vals)-1)
}

func inOrderTraversal(root *TreeNode) []int {
	ret := make([]int,0)
	for root != nil {
		if root.Left == nil {
			ret = append(ret, root.Val)
			root = root.Right
		}else if root.Left != nil {
			pivot := root.Left
			for pivot.Right != nil && pivot.Right != root {
				pivot = pivot.Right
			}
			if pivot.Right == nil {
				pivot.Right = root
				root = root.Left
			}else if pivot.Right == root {
				pivot.Right = nil
				ret = append(ret, root.Val)
				root = root.Right
			}
		}
	}
	return ret
}

func buildTree(vals []int, left int, right int) *TreeNode {
	if len(vals) == 0 {
		return nil
	}
	if left > right {
		return nil
	}
	if left == right {
		return &TreeNode{
			Val:   vals[left],
		}
	}
	mid := left + (right - left) >> 1
	return &TreeNode{
		Val:   vals[mid],
		Left:  buildTree(vals, left, mid -1),
		Right: buildTree(vals, mid+1, right),
	}
}

func TestBalanceBST(t *testing.T) {
	root := &TreeNode{
		Val:   1,
		Right: &TreeNode{
			Val:   2,
			Right:  &TreeNode{
				Val:   3,
				Right: &TreeNode{
					Val:   4,
				},
			},
		},
	}

	bst := balanceBST(root)
	t.Log(bst)
}