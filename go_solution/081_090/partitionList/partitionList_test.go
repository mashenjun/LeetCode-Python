package partitionList

import (
	"testing"
)

/*
	86. Partition List
	Given a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.
	You should preserve the original relative order of the nodes in each of the two partitions.
 */

 type ListNode struct {
     Val int
     Next *ListNode
 }

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyBefore := &ListNode{
		Val:  0,
		Next: nil,
	}
	before := dummyBefore
	dummyAfter := &ListNode{
		Val:  0,
		Next: nil,
	}
	after := dummyAfter
	for head != nil {
		if head.Val < x {
			before.Next = head
			before = before.Next
			if head.Next == nil {
				// clear after chain
				after.Next = nil
			}
		}else {
			after.Next = head
			after = after.Next
			if head.Next == nil {
				// clear before chain
				before.Next = nil
			}
		}
		head = head.Next
	}
	before.Next = dummyAfter.Next
	return dummyBefore.Next
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

func Test_partition(t *testing.T) {
	head := genLinkedList([]int{1,4,3,2,5,2})
	rlt := partition(head, 3)
	t.Log(partition(head, 3))
	for rlt != nil{
		t.Log(rlt.Val)
		rlt = rlt.Next
	}
	//output: 1->2->2->4->3->5
}