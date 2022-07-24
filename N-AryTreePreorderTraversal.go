package leetcode

func preorder(root *Node) []int {

	if root == nil {
		return []int{}
	}

	var out = []int{root.Val}

	if root.Children != nil {
		for _, child := range root.Children {
			out = append(out, preorder(child)...)
		}
	}

	return out
}
