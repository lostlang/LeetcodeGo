package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func generateListNode(array []int) *ListNode {
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

func (node ListNode) New(array []int) *ListNode {
	return generateListNode(array)
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

func (node TreeNode) New(array []interface{}) *TreeNode {
	return generateTreeNode(array)
}

type Node struct {
	Val      int
	Children []*Node
}

func generateNode(array []interface{}) *Node {
	if len(array) == 0 {
		return nil
	}

	outNode := &Node{array[0].(int), nil}
	array = array[1:]

	currentLevel := []*Node{outNode}
	nextLevel := []*Node{outNode}
	var newNode *Node

	for _, val := range array {
		if val == nil {
			currentLevel = currentLevel[1:]

			if len(currentLevel) == 0 {
				currentLevel = nextLevel
				nextLevel = []*Node{}
			}
		} else {
			newNode = &Node{val.(int), nil}
			if currentLevel[0].Children == nil {
				currentLevel[0].Children = []*Node{}
			}
			currentLevel[0].Children = append(currentLevel[0].Children, newNode)
			nextLevel = append(nextLevel, newNode)
		}
	}

	return outNode
}

func (node Node) New(array []interface{}) *Node {
	return generateNode(array)
}

func printListNode(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
	}
	fmt.Println()
}
