package leetcode

func postorder(root *Node) []int {

	if root == nil {
		return []int{}
	}

	var out []int

	if root.Children != nil {
		for _, child := range root.Children {
			out = append(out, postorder(child)...)
		}
	}

	out = append(out, root.Val)

	return out
}
