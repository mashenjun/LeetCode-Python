package removeNthNodeFromEndOfList

type ListNode struct {
	Val  int
	Next *ListNode
}
// 19. Remove Nth Node From End of List
// 思路：利用一前一后两个指针
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	early, later := head, head
	for i := 0; i < n; i++ {
		early = early.Next
	}
	var prev *ListNode
	for early != nil {
		early = early.Next
		prev, later = later, later.Next
	}
	if prev == nil {
		return later.Next
	}
	prev.Next = later.Next
	return head
}
