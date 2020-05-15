package pointOffer

// 定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。

// 思路：利用额外的一个指针记录前继节点
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		prev, head, head.Next = head, head.Next, prev
	}
	return prev
}
