package leetcode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	tmpA := headA
	tmpB := headB
	var len int

	for tmpA != nil {
		tmpA = tmpA.Next
		len++
	}

	for tmpB != nil {
		tmpB = tmpB.Next
		len--
	}

	if len > 0 {
		tmpA = headA
		tmpB = headB
	} else {
		tmpA = headB
		tmpB = headA
		len = -len
	}

	for i := 0; i < len; i++ {
		tmpA = tmpA.Next
	}

	for tmpA != tmpB {
		tmpA = tmpA.Next
		tmpB = tmpB.Next
	}

	return tmpA
}
