package addTwoNumbersII

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1.Val == 0 {
		return l2
	} else if l2.Val == 0 {
		return l1
	}
	r1 := reverse(l1)
	r2 := reverse(l2)

	var prev *ListNode
	remain, val := 0, 0
	for r1 != nil && r2 != nil {
		val, remain = (r1.Val+r2.Val+remain)%10, (r1.Val+r2.Val+remain)/10
		node := &ListNode{
			Val:  val,
			Next: prev,
		}
		prev = node
		r1, r2 = r1.Next, r2.Next
	}
	for r1 != nil {
		val, remain = (r1.Val+remain)%10, (r1.Val+remain)/10
		node := &ListNode{
			Val:  val,
			Next: prev,
		}
		prev = node
		r1 = r1.Next
	}
	for r2 != nil {
		val, remain = (r2.Val+remain)%10, (r2.Val+remain)/10
		node := &ListNode{
			Val:  val,
			Next: prev,
		}
		prev = node
		r2 = r2.Next
	}
	return prev
}

func reverse(root *ListNode) *ListNode {
	if root == nil {
		return root
	}
	var prev *ListNode
	curr := root
	for curr != nil {
		curr, curr.Next, prev = curr.Next, prev, curr
	}
	return prev
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{
		Val:  7,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  4,
				Next: &ListNode{
					Val:  3,
				},
			},
		},
	}

	l2 := &ListNode{
		Val:  5,
		Next: &ListNode{
			Val:  6,
			Next: &ListNode{
				Val:  4,
			},
		},
	}

	head := addTwoNumbers(l1, l2)
	t.Log(head)
}
