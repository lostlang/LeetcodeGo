package utils

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(array []interface{}) *TreeNode {
	if len(array) == 0 {
		return nil
	}

	outTree := &TreeNode{array[0].(int), nil, nil}

	currentLevel := []*TreeNode{outTree}
	nextLevel := []*TreeNode{}
	var newNode *TreeNode

	array = array[1:]

	for index, elem := range array {
		if len(currentLevel) == 0 {
			currentLevel = nextLevel
			nextLevel = []*TreeNode{}
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
