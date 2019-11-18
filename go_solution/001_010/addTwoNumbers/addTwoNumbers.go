package addTwoNumbers

/*
	002. Add Two Numbers
	Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
	Output: 7 -> 0 -> 8
	Explanation: 342 + 465 = 807.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	tmp := 0 // 记录进位
	dummy := &ListNode{Val: 0, Next: nil} // dummy node to track the result
	node := dummy
	for l1 != nil || l2 != nil || tmp != 0 {
		if l1 != nil {
			tmp = tmp + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp = tmp + l2.Val
			l2 = l2.Next
		}
		node.Next = &ListNode{
			Val:  tmp % 10,
			Next: nil,
		}
		node = node.Next
		tmp = tmp / 10
	}
	return dummy.Next
}
