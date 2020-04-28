package binaryTreeInorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	return inorderIteration(root)
}
// 非递归 中序遍历
// 思路：先将从curr开始把所有可以到达的左子入栈
// 		然后curr=stack.Pop()，并将curr指向自己的右子，再从curr开始把所有可以到达的左子入栈然后再消费栈顶
func inorderIteration(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var (
		// 需要 stack 和一个 curr指针 来控制访问的顺序
		stack = []*TreeNode{}
		curr  = root
		rlt   = make([]int, 0)
	)
	for curr != nil || len(stack) != 0 {
		// 先把从curr开始所有需要访问的左子入栈
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		// 出栈并访问
		curr, stack =  stack[len(stack)-1], stack[:len(stack)-1]
		rlt = append(rlt, curr.Val)
		// curr 指向右子
		curr = curr.Right
	}
	return rlt
}

// 递归 中序遍历
func inorderRecursive(root *TreeNode) []int {
	if root != nil {
		rlt := inorderRecursive(root.Left)
		rlt = append(rlt, root.Val)
		return append(rlt, inorderRecursive(root.Right)...)
	} else {
		return []int{}
	}
}

// 1. 如果当前节点的左孩子为空，则输出当前节点并将其右孩子作为当前节点。
// 2. 如果当前节点的左孩子不为空，在当前节点的左子树中找到当前节点在中序遍历下的前驱节点。
//  a) 如果前驱节点的右孩子为空，将它的右孩子设置为当前节点。当前节点更新为当前节点的左孩子。
//  b) 如果前驱节点的右孩子为当前节点，将它的右孩子重新设为空（恢复树的形状）,输出当前节点。当前节点更新为当前节点的右孩子。
// 3. 重复以上1、2直到当前节点为空。

func morrisTraversal(root *TreeNode) []int {
	ret := make([]int,0)
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
				root = root.Left
			}else if pivot.Right == root{
				pivot.Right = nil
				ret = append(ret, root.Val)
				root = root.Right
			}
		}
	}
	return ret
}

