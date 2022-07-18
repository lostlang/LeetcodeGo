package leetcode

import (
	"reflect"
	"testing"
)

type countOperationsTestPair struct {
	input1 int
	input2 int
	out    int
}

var countOperationsTestCases = []countOperationsTestPair{
	{2, 3, 3},
	{10, 10, 1},
	{0, 0, 0},
	{0, 1, 0},
}

func TestEvalCountOperations(t *testing.T) {
	for _, pair := range countOperationsTestCases {
		var newOut = countOperations(pair.input1, pair.input2)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input1,
				"and", pair.input2,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
