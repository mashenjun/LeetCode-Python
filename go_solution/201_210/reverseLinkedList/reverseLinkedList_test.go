package reverseLinkedList

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		// 三个赋值语句同时执行
		head, head.Next, prev = head.Next, prev, head
	}
	return prev
}

func TestReverseList(t *testing.T) {
	head := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	newHead := reverseList(head)
	t.Log(newHead)
}