package addTwoNumbers

import (
	"strconv"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := ListNode{
		Val:  2,
		Next: &ListNode{
			Val:  4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := ListNode{
		Val:  5,
		Next: &ListNode{
			Val:  6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	ret := ""
	head := addTwoNumbers(&l1, &l2)
	for head != nil {
		ret = ret+ strconv.Itoa(head.Val)
		head = head.Next
	}
	if ret != "708" {
		t.Fatalf("error result:%+v", ret)
	}
}