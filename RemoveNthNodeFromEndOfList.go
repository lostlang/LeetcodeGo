package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	var arr []int

	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	arr = append(arr[:len(arr)-n], arr[len(arr)-n+1:]...)

	return generateListNode(arr)
}
