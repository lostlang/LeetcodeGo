package utils

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(array []int) *ListNode {
	if len(array) == 0 {
		return nil
	}

	outList := &ListNode{array[0], nil}

	array = array[1:]

	curList := outList

	for _, elem := range array {
		curList.Next = &ListNode{elem, nil}
		curList = curList.Next
	}

	return outList
}
