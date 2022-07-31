package leetcode

func getDecimalValue(head *ListNode) int {
	var out int

	for head != nil {
		out *= 2
		out += head.Val
		head = head.Next
	}

	return out
}
