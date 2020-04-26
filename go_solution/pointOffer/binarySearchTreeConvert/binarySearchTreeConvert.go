package binarySearchTreeConvert

/*
题目描述
输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的双向链表。要求不能创建任何新的结点，只能调整树中结点指针的指向。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Prev *ListNode
	Next *ListNode
}

func convert(root *TreeNode) *ListNode {
	if root == nil {
		return nil
	}
	head, _ := inorderIteration(root)
	return head
}

// 方法一：中序遍历， 递归
func inorderRecursive(root *TreeNode) (head *ListNode, tail *ListNode) {
	if root == nil {
		return nil, nil
	}
	curr := &ListNode{
		Val: root.Val,
	}
	head, tail = curr, curr
	lHead, lTail := inorderRecursive(root.Left)
	if lHead != nil && lTail != nil {
		lTail.Next, curr.Prev = curr, lTail
		head = lHead
	}
	rHead, rTail := inorderRecursive(root.Right)
	if rHead != nil && rTail != nil {
		curr.Next, rHead.Prev = rHead, curr
		tail = rTail
	}
	return
}

// 方法二：中序遍历， 非递归
func inorderIteration(root *TreeNode) (head *ListNode, tail *ListNode) {
	var (
		stack = make([]*TreeNode, 0)
		curr  = root
		prev  *ListNode
	)

	for len(stack) != 0 || curr != nil {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr, stack = stack[len(stack)-1], stack[:len(stack)-1]
		listNode := &ListNode{Val: curr.Val}
		if head == nil {
			head = listNode
		} else {
			prev.Next, listNode.Prev = listNode, prev
		}
		prev = listNode
		curr = curr.Right
	}
	tail = prev
	return head, tail
}
