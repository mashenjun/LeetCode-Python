package reorderList

import "testing"

// 143. Reorder List

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归+单指针，速度非常慢
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	last, next := head, head.Next
	prev := &ListNode{
		Next: head,
	}
	for last.Next != nil {
		prev, last = last, last.Next
	}
	if head.Next == last {
		return
	}
	head.Next, last.Next, prev.Next = last, head.Next, nil
	reorderList(next)
}

func TestReorderList(t *testing.T) {
	// [1,2,3,4]
	head := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  4,
				},
			},
		},
	}
	reorderList(head)
}
