package convertSortedListToBinarySearchTree

// 109. Convert Sorted List to Binary Search Tree
// Given a singly linked list where elements are sorted in ascending order, convert it to a height balanced BST.
// For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.
type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	return toTree(head)
}
// 递归实现，每次递归都需要用快慢指针找到链表的中间节点，然后拆分成左半部分，中间节点，右半部分
// 利用左半部分，中间节点，右半部分构造成一个BST是
func toTree(head *ListNode) (root *TreeNode) {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{
			Val:   head.Val,
			Left:  nil,
			Right: nil,
		}
	}
	// find middle node in the list
	slow, fast, slowPrev := head, head, &ListNode{Next:head}
	for fast != nil && fast.Next != nil {
		slow, slowPrev = slow.Next, slow
		fast = fast.Next.Next
	}
	// slow now points to the middle node
	slowPrev.Next = nil
	return &TreeNode{
		Val:   slow.Val,
		Left:  toTree(head),
		Right: toTree(slow.Next),
	}
}
