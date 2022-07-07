package leetcode

func generateListNode(array []int) *ListNode {
	var outList = &ListNode{array[0], nil}

	array = array[1:]

	for _, elem := range array {
		outList = &ListNode{elem, outList}
	}
	return outList
}
