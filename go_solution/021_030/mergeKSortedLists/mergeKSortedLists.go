package mergeKSortedLists

// 023. Merge k Sorted Lists
// Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

type ListNode struct {
	Val  int
	Next *ListNode
}

// naive approach
// do merge two list k-1 times
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}else if len(lists) == 1 {
		return lists[0]
	}
	head := mergeTwoLists(lists[0], lists[1])
	for i:= 2; i< len(lists); i++ {
		head = mergeTwoLists(head, lists[i])
	}
	return head
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	dummy := &ListNode{}
	head := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			l1, head = l1.Next, head.Next
		}else {
			head.Next = l2
			l2, head = l2.Next, head.Next
		}
	}
	if l1 != nil {
		head.Next = l1
	}
	if l2 != nil {
		head.Next = l2
	}
	return dummy.Next
}
