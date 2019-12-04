package reverseLinkedListII

import "testing"

/*
	Reverse a linked list from position m to n. Do it in one-pass.
	Note: 1 ≤ m ≤ n ≤ length of list.

	Input: 1->2->3->4->5->NULL, m = 2, n = 4
	Output: 1->4->3->2->5->NULL
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head.Next == nil || m == n {
		return head
	}
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	var prev *ListNode
	var pivot, tail *ListNode
	for i:=1;i<m;i++ {
		prev, head = head, head.Next
	}
	pivot, tail = prev, head
	var tmp *ListNode
	for i:=0;i<=n-m;i++ {
		tmp = head.Next
		head.Next, prev = prev, head
		head = tmp
	}
	// join the linked list
	tail.Next = head
	if pivot != nil {
		pivot.Next = prev
	}else {
		dummy.Next = prev
	}
	return dummy.Next
}

func genLinkedList(source []int) *ListNode {
	if len(source) == 0 {
		return nil
	}
	dummy := &ListNode{
		Val:  0,
		Next: nil,
	}
	curr := dummy
	for _, v := range source {
		curr.Next = &ListNode{
			Val:  v,
			Next: nil,
		}
		curr = curr.Next
	}
	return dummy.Next
}

func Test_reverseBetween(t *testing.T) {
	head := genLinkedList([]int{1,2,3,4,5})
	rlt := reverseBetween(head,2,4)
	for rlt != nil {
		t.Log(rlt.Val)
		rlt = rlt.Next
	}
}
