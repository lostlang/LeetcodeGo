package leetcode

func hasCycle(head *ListNode) bool {
	var hash = make(map[*ListNode]bool)

	for head != nil {
		if hash[head] {
			return true
		} else {
			hash[head] = true
		}

		head = head.Next
	}

	return false
}
