package splitLinkedListInParts

import (
	"fmt"
	"testing"
)

// 725. Split Linked List in Parts
// Given a (singly) linked list with head node root, write a function to split the linked list into k consecutive linked list "parts".

type ListNode struct {
	Val int
	Next *ListNode
}

// 首先遍历一边链表来获取链表的长度，然后根据长度和k，计算出每个part的分布情况
// 根据计算出的分布情况，再次遍历链表，构造ret
func splitListToParts(root *ListNode, k int) []*ListNode {
	total := 0
	curr := root
	for curr != nil {
		curr = curr.Next
		total ++
	}
	ret := make([]*ListNode, k)
	distribution := calDistribution(total, k)
	prev := &ListNode{Next:root}
	for i:= 0; i<len(distribution); i++ {
		if distribution[i] == 0 {
			break
		}
		ret[i] = root
		for distribution[i] > 0 {
			prev, root = root, root.Next
			distribution[i]--
			if distribution[i] == 0 {
				prev.Next = nil
			}
		}
	}
	return ret
}

func calDistribution(total, k int) []int {
	distribution := make([]int, k)
	part := total / k
	if part > 0 {
		for i:= 0; i< len(distribution); i++ {
			distribution[i] = part
		}
	}
	remain := total % k
	for i:=0;i<remain;i++ {
		distribution[i]++
	}
	return distribution
}

func Test_splitListToParts(t *testing.T) {
	input := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}

	ret := splitListToParts(input, 5)
	fmt.Println(ret)
}