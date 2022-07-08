package leetcode

func generateListNode(array []int) *ListNode {
	if len(array) == 0 {
		return nil
	}

	var outList = &ListNode{array[0], nil}

	array = array[1:]

	var curList = outList

	for _, elem := range array {
		curList.Next = &ListNode{elem, nil}
		curList = curList.Next
	}

	return outList
}
