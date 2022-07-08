package leetcode

import (
	"reflect"
	"testing"
)

type distributeCandiesTestPair struct {
	inputCandy  int
	inputPeople int
	out         []int
}

var distributeCandiesTestCases = []distributeCandiesTestPair{
	{7, 4, []int{1, 2, 3, 1}},
	{10, 3, []int{5, 2, 3}},
}

func TestEvalDistributeCandies(t *testing.T) {
	for _, pair := range distributeCandiesTestCases {
		var newOut = distributeCandies(pair.inputCandy, pair.inputPeople)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputCandy,
				"and", pair.inputPeople,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
