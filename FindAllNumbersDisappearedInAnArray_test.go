package leetcode

import (
	"reflect"
	"testing"
)

type findDisappearedNumbersTestPair struct {
	input []int
	out   []int
}

var findDisappearedNumbersTestCases = []findDisappearedNumbersTestPair{
	{[]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{5, 6}},
	{[]int{1, 1}, []int{2}},
}

func TestEvalFindDisappearedNumbers(t *testing.T) {
	for _, pair := range findDisappearedNumbersTestCases {
		var newOut = findDisappearedNumbers(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
