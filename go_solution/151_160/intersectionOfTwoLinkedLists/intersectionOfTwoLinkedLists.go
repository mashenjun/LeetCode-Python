package intersectionOfTwoLinkedLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	currA, currB := headA, headB
	countA, countB := 0, 0
	for currA != nil {
		currA = currA.Next
		countA++
	}
	for currB != nil {
		currB = currB.Next
		countB++
	}
	if countA > countB {
		for i := 0; i < countA-countB; i++ {
			headA = headA.Next
		}
	} else if countA < countB {
		for i := 0; i < countB-countA; i++ {
			headB = headB.Next
		}
	}
	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}
