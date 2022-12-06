package leetcode

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var cur = head
	var length = 0
	for cur != nil {
		cur = cur.Next
		length++
	}
	k = k % length
	k = length - k

	if k == 0 {
		return head
	}

	cur = head.Next
	var copyHead = &ListNode{head.Val, nil}
	var copyHeadStart = copyHead
	k--

	for ; k > 0; k-- {
		var newHead = &ListNode{cur.Val, nil}
		copyHead.Next = newHead
		copyHead = newHead
		cur = cur.Next
	}

	var out = cur

	if cur == nil {
		return copyHeadStart
	}

	for cur.Next != nil {
		cur = cur.Next
	}

	cur.Next = copyHeadStart

	return out
}
