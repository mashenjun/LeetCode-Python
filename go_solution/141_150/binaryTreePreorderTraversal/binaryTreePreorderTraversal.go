package binaryTreePreorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	return preorderInteration(root)
}
// 非递归 前序遍历
func preorderInteration(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var (
		// 利用 stack 来存储需要需要访问的节点
		stack = []*TreeNode{root}
		rlt   = make([]int, 0)
	)
	for len(stack) != 0 {
		// 出栈，直接访问
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		rlt = append(rlt, curr.Val)
		// 先将右子入栈，再将左子入栈
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
	}
	return rlt
}

// 递归 前序遍历
func preorderRecursive(root *TreeNode) []int {
	if root == nil {
		return []int{}
	} else {
		rlt := []int{root.Val}
		rlt = append(rlt, preorderRecursive(root.Left)...)
		rlt = append(rlt, preorderRecursive(root.Right)...)
		return rlt
	}
}

// morris 前序遍历
// 1. 如果当前节点的左孩子为空，则输出当前节点并将其右孩子作为当前节点。
// 2. 如果当前节点的左孩子不为空，在当前节点的左子树中找到当前节点在中序遍历下的前驱节点。
//  a) 如果前驱节点的右孩子为空，将它的右孩子设置为当前节点。输出当前节点，当前节点更新为当前节点的左孩子。
//  b) 如果前驱节点的右孩子为当前节点，将它的右孩子重新设为空（恢复树的形状）。当前节点更新为当前节点的右孩子。
// 3. 重复以上1、2直到当前节点为空。
func preorderMorris(root *TreeNode) []int {
	ret := make([]int, 0)
	for root != nil {
		if root.Left == nil {
			ret = append(ret, root.Val)
			root = root.Right
		}else {
			pivot := root.Left
			for pivot.Right != nil && pivot.Right != root {
				pivot = pivot.Right
			}
			if pivot.Right == nil {
				pivot.Right = root
				ret = append(ret, root.Val)
				root = root.Left
			}else if pivot.Right == root{
				pivot.Right = nil
				root = root.Right
			}
		}
	}
	return ret
}