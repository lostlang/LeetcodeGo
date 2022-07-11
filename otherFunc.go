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

func generateTreeNode(array []interface{}) *TreeNode {

	if len(array) == 0 {
		return nil
	}

	var outTree = &TreeNode{array[0].(int), nil, nil}

	var currentLevel = []*TreeNode{outTree}
	var nextLevel = []*TreeNode{}
	var newNode *TreeNode

	array = array[1:]

	for index, elem := range array {

		if len(currentLevel) == 0 {
			currentLevel = nextLevel
		}

		if elem != nil {
			newNode = &TreeNode{elem.(int), nil, nil}
			nextLevel = append(nextLevel, newNode)
		} else {
			newNode = nil
		}

		if index%2 == 0 {
			currentLevel[0].Left = newNode
		} else {
			currentLevel[0].Right = newNode
			currentLevel = currentLevel[1:]
		}

	}

	return outTree
}
