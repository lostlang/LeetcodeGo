package leetcode

import (
	"reflect"
	"testing"
)

type numFactoredBinaryTreesTestPair struct {
	input []int
	out   int
}

var numFactoredBinaryTreesTestCases = []numFactoredBinaryTreesTestPair{
	{[]int{2, 4}, 3},
	{[]int{2, 4, 5, 10}, 7},
}

func TestEvalNumFactoredBinaryTrees(t *testing.T) {
	for _, pair := range numFactoredBinaryTreesTestCases {
		var newOut = numFactoredBinaryTrees(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
