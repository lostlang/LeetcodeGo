package leetcode

func removeElements(head *ListNode, val int) *ListNode {
	var outList *ListNode
	if head == nil {
		return outList
	}

	for outList == nil && head != nil {
		if head.Val != val {
			outList = &ListNode{head.Val, nil}
		}
		head = head.Next
	}

	var curList = outList

	for head != nil {
		if head.Val != val {
			curList.Next = &ListNode{head.Val, nil}
			curList = curList.Next
		}
		head = head.Next
	}

	return outList
}
