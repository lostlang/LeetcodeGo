package leetcode

import (
	"reflect"
	"testing"
)

type kidsWithCandiesTestPair struct {
	inputCandies      []int
	inputExtraCandies int
	out               []bool
}

var kidsWithCandiesTestCases = []kidsWithCandiesTestPair{
	{[]int{2, 3, 5, 1, 3}, 3, []bool{true, true, true, false, true}},
	{[]int{4, 2, 1, 1, 2}, 1, []bool{true, false, false, false, false}},
	{[]int{12, 1, 12}, 10, []bool{true, false, true}},
}

func TestEvalKidsWithCandies(t *testing.T) {
	for _, pair := range kidsWithCandiesTestCases {
		var newOut = kidsWithCandies(pair.inputCandies, pair.inputExtraCandies)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputCandies,
				"and", pair.inputExtraCandies,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
