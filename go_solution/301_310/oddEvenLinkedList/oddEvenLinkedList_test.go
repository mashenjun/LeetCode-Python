package oddEvenLinkedList

import "testing"

// 328. Odd Even Linked List

// Given a singly linked list, group all odd nodes together followed by the even nodes.
// Please note here we are talking about the node number and not the value in the nodes.
//You should try to do it in place. The program should run in O(1) space complexity and O(nodes) time complexity.

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	odd, even := head, head.Next // 用两个指针去跟踪odd位置的节点和even位置的节点
	evenHead := even
	for even != nil && even.Next != nil {
		even.Next, odd.Next = even.Next.Next, odd.Next.Next // 一定要同时赋值
		even, odd = even.Next, odd.Next                     // 一定要同时赋值
	}
	odd.Next = evenHead // odd和even做拼接
	return head
}

func Test_OddEvenList(t *testing.T) {
	input := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	_ = oddEvenList(input)

}
