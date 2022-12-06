package leetcode

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var odd = &ListNode{head.Val, nil}
	var oddHead = odd
	if head.Next == nil {
		return odd
	}
	var even = &ListNode{head.Next.Val, nil}
	var evenHead = even

	var cur = head.Next.Next

	var flag = true

	for cur != nil {
		if flag {
			odd.Next = &ListNode{cur.Val, nil}
			odd = odd.Next
		} else {
			even.Next = &ListNode{cur.Val, nil}
			even = even.Next
		}
		flag = !flag
		cur = cur.Next
	}

	odd.Next = evenHead

	return oddHead
}
