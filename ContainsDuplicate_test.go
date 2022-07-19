package leetcode

import (
	"reflect"
	"testing"
)

type containsDuplicateTestPair struct {
	input []int
	out   bool
}

var containsDuplicateTestCases = []containsDuplicateTestPair{
	{[]int{1, 2, 3, 1}, true},
	{[]int{1, 2, 3, 4}, false},
	{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
}

func TestEvalContainsDuplicate(t *testing.T) {
	for _, pair := range containsDuplicateTestCases {
		var newOut = containsDuplicate(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
