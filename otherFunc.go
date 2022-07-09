package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTreeNode(array []int) *TreeNode {

	if len(array) == 0 {
		return nil
	}

	var outTree = &TreeNode{array[0], nil, nil}

	array = array[1:]
	// Face realization
	for index, elem := range array {
		if index%2 == 0 {
			outTree.Left = &TreeNode{elem, nil, nil}
		} else {
			outTree.Right = &TreeNode{elem, nil, nil}
		}

	}
	return outTree
}
