package leetcode

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
