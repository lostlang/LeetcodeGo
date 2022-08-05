package leetcode

func detectCycle(head *ListNode) *ListNode {
	var listHead = make([]*ListNode, 0)

	for head != nil {
		for _, val := range listHead {
			if val == head.Next {
				return val
			}
		}

		listHead = append(listHead, head)
		head = head.Next
	}

	return nil
}
