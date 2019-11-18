package removeDuplicatesFromSortedListII

import "testing"

/*
	82. Remove Duplicates from Sorted List
	Given a sorted linked list, delete all nodes that have duplicate numbers,
	leaving only distinct numbers from the original list.
	IDEA: 由于是排序好的链表，可以直接在遍历的时候判断
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// define dummy node to track head
	dummy := &ListNode{
		Next: head,
	}
	curr := dummy
	for curr.Next != nil{
		if curr.Next.Next != nil && curr.Next.Val == curr.Next.Next.Val {
			value :=curr.Next.Val
			for curr.Next != nil && curr.Next.Val == value {
				curr.Next = curr.Next.Next
			}
		} else {
			curr = curr.Next
		}
	}
	return dummy.Next
}

func Test_deleteDuplicates(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next:&ListNode{
						Val: 3,
						Next: nil,
					},
				},
			},
		},
	}
	res := deleteDuplicates(head)
	for res != nil {
		t.Log(res.Val)
		res = res.Next
	}
}