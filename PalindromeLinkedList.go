package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var array []int

	for head != nil {
		array = append(array, head.Val)
		head = head.Next
	}

	for index := range array {
		if array[index] != array[len(array)-index-1] {
			return false
		}
	}

	return true
}

func generateListNode(array []int) *ListNode {
	var outList = &ListNode{array[0], nil}

	array = array[1:]

	for _, elem := range array {
		outList = &ListNode{elem, outList}
	}
	return outList
}
