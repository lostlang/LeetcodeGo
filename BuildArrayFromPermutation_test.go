package leetcode

import (
	"reflect"
	"testing"
)

type buildArrayTestPair struct {
	input []int
	out   []int
}

var buildArrayTestCases = []buildArrayTestPair{
	{[]int{0, 2, 1, 5, 3, 4}, []int{0, 1, 2, 4, 5, 3}},
	{[]int{5, 0, 1, 2, 3, 4}, []int{4, 5, 0, 1, 2, 3}},
}

func TestEvalBuildArray(t *testing.T) {
	for _, pair := range buildArrayTestCases {
		var newOut = buildArray(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
