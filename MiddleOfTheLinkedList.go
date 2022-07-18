package leetcode

func middleNode(head *ListNode) *ListNode {
	var cur = head
	var counter int

	for cur != nil {
		cur = cur.Next
		counter++
	}

	counter /= 2

	for i := 0; i < counter; i++ {
		head = head.Next
	}

	return head
}
