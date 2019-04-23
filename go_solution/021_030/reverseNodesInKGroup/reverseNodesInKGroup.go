package reverseNodesInKGroup

type ListNode struct {
	Val  int
	Next *ListNode
}

/* 025. Reverse Nodes in k-Group
Given this linked list: 1->2->3->4->5
For k = 2, you should return: 2->1->4->3->5
For k = 3, you should return: 3->2->1->4->5
*/
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil { // no need to reverse
		return head
	}
	if k < 2 { // no need to reverse
		return head
	}
	var rlt *ListNode
	nextHead := head
	// 如果k大于整条链的长度，也不用反转，直接返回head
	for i := k; i > 1; i-- {
		nextHead = nextHead.Next
		if nextHead == nil {
			return head
		}
	}
	// 如果k小于整条链的长度，此时nextHead指向改group中反转后的头节点
	// 需要反转
	rlt = nextHead
	current := head
	for nextHead != nil {
		tail := current
		var perv *ListNode = nil
		for j := 0; j < k; j++ {
			next := current.Next
			current.Next = perv
			perv = current
			current = next
			if nextHead != nil {
				nextHead = nextHead.Next
			}
		}
		if nextHead != nil {
			tail.Next = nextHead
		} else {
			tail.Next = current
		}
	}
	return rlt
}
