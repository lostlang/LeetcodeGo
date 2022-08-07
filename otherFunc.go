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

	var outTree = &TreeNode{array[0].(int), nil, nil}

	var currentLevel = []*TreeNode{outTree}
	var nextLevel = []*TreeNode{}
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

	var outNode = &Node{array[0].(int), nil}
	array = array[1:]

	var currentLevel = []*Node{outNode}
	var nextLevel = []*Node{outNode}
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

func stringToHashmapByte(str string) map[byte]int {
	var hash = make(map[byte]int)

	for i := range str {
		hash[str[i]]++
	}

	return hash
}

func stringToHashmapRune(str string) map[rune]int {
	var hash = make(map[rune]int)

	for _, val := range str {
		hash[val]++
	}

	return hash
}

func intToHash(arr []int) map[int]int {
	var out = make(map[int]int)

	for _, val := range arr {
		out[val]++
	}

	return out
}

func stringToArrayInts(s string) []int {
	var out []int

	for _, char := range s {
		out = append(out, int(char-'0'))
	}

	return out
}

func reverseInts(a []int) []int {
	var out []int

	for i := len(a) - 1; i >= 0; i-- {
		out = append(out, a[i])
	}

	return out
}
