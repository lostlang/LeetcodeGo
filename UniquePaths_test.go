package leetcode

import (
	"reflect"
	"testing"
)

type uniquePathsTestPair struct {
	inputM int
	inputN int
	out    int
}

var uniquePathsTestCases = []uniquePathsTestPair{
	{3, 7, 28},
	{3, 2, 3},
}

func TestEvalUniquePaths(t *testing.T) {
	for _, pair := range uniquePathsTestCases {
		var newOut = uniquePaths(pair.inputM, pair.inputN)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputM,
				"and", pair.inputN,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
