package leetcode

import (
	"reflect"
	"testing"
)

func TestEvalLowestCommonAncestorTree1(t *testing.T) {
	root := generateTreeNode([]interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5})
	q := root.Left
	p := root.Right
	out := root

	var newOut = lowestCommonAncestor(root, p, q)
	if !reflect.DeepEqual(newOut, out) {
		t.Error(
			"For", root, p, q,
			"expected", out,
			"got", newOut,
		)
	}
}

func TestEvalLowestCommonAncestorTree2(t *testing.T) {
	root := generateTreeNode([]interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5})
	q := root.Left
	p := root.Left.Left
	out := root.Left

	var newOut = lowestCommonAncestor(root, p, q)
	if !reflect.DeepEqual(newOut, out) {
		t.Error(
			"For", root, p, q,
			"expected", out,
			"got", newOut,
		)
	}
}
