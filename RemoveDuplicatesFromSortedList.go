package leetcode

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var out = head
	var hash = make(map[int]bool)
	hash[head.Val] = true

	var cur = out

	for head != nil {
		if !hash[head.Val] {
			cur.Next = head
			hash[head.Val] = true
			cur = cur.Next
		}

		head = head.Next
	}

	cur.Next = nil

	return out
}
