package removeDuplicatesFromSortedList

// 83. Remove Duplicates from Sorted List
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	existed := make(map[int]struct{})
	dummy := ListNode{
		Next: head,
	}
	var prev *ListNode
	for head != nil {
		if _, ok := existed[head.Val]; !ok {
			existed[head.Val] = struct{}{}
			prev, head = head, head.Next
		} else {
			prev.Next, head = head.Next, head.Next
		}
	}
	return dummy.Next
}
