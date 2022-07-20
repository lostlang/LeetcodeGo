package leetcode

func reverseList(head *ListNode) *ListNode {
	var array []int

	for head != nil {
		array = append(array, head.Val)
		head = head.Next
	}

	for i := 0; i < len(array)/2; i++ {
		array[i], array[len(array)-1-i] = array[len(array)-1-i], array[i]
	}

	return generateListNode(array)
}
