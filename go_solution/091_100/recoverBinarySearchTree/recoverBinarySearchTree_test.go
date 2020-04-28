package recoverBinarySearchTree

import "testing"

// 99. Recover Binary Search Tree
// Two elements of a binary search tree (BST) are swapped by mistake.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func recoverTree(root *TreeNode)  {

}

// 利用morris中序遍历整个tree，如果出现非第增的情况，就是我们需要交换的节点
// 因为我们有两个节点是在错误的位置，在中序遍历的时候会有两次非递增的情况出现
// 第一次非递增的情况 a.val > b.val; a 是我们要找的第一个错误节点
// 第二次非递增的情况 x.val > y.val; y 是我们要找的第二个错误节点
// 然后交换 a.val 和 y.val 或者交换 a 节点和 y 节点来修复BST
func recoverWithMorris(root *TreeNode)  {
	var wrong1 *TreeNode = nil
	var wrong2 *TreeNode = nil
	var prev  *TreeNode = nil

	for root != nil {
		if root.Left == nil {
			if prev != nil && prev.Val > root.Val{
				if wrong1 != nil {
					wrong2 = root
				}else {
					wrong1, wrong2 = prev, root
				}
			}
			prev, root = root, root.Right
		}else {
			pivot := root.Left
			for pivot.Right !=nil && pivot.Right != root {
				pivot = pivot.Right
			}
			if pivot.Right == nil {
				pivot.Right = root
				root = root.Left
			}else if pivot.Right == root {
				pivot.Right = nil

				if prev != nil && prev.Val > root.Val{
					if wrong1 != nil {
						wrong2 = root
					}else {
						wrong1, wrong2 = prev, root
					}
				}
				prev, root = root, root.Right
			}
		}
	}
	swap(wrong1, wrong2)
}

func swap(a *TreeNode, b *TreeNode) {
	a.Val, b.Val = b.Val, a.Val
}

// [3,1,4,null,null,2]
func Test_RecoverWithMorris(t *testing.T) {
	root := &TreeNode{
		Val:   3,
		Left:  &TreeNode{
			Val:   1,
		},
		Right: &TreeNode{
			Val:   4,
			Left:  &TreeNode{
				Val:   2,
			},
		},
	}
	recoverWithMorris(root)
}
