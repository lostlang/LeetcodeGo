package leetcode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var out = &ListNode{}
	var current = out
	var flag = false

	for l1 != nil || l2 != nil {
		var sum = 0

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if flag {
			sum++
			flag = false
		}

		if sum >= 10 {
			sum -= 10
			flag = true
		}

		current.Val = sum
		current.Next = &ListNode{}
		current = current.Next
	}

	for l1 != nil {
		var val = l1.Val
		if flag {
			val++
			flag = false
		}

		if val >= 10 {
			val -= 10
			flag = true
		}

		current.Val = val
		current.Next = &ListNode{}
		current = current.Next
		l1 = l1.Next
	}

	for l2 != nil {
		var val = l2.Val
		if flag {
			val++
			flag = false
		}

		if val >= 10 {
			val -= 10
			flag = true
		}

		current.Val = val
		current.Next = &ListNode{}
		current = current.Next
		l2 = l2.Next
	}

	if flag {
		current.Val = 1
	}

	var cur = out
	for cur.Next != nil {
		if cur.Next.Next == nil && cur.Next.Val == 0 {
			cur.Next = nil
			break
		}
		cur = cur.Next
	}

	return out
}
