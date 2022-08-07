package leetcode

import (
	"reflect"
	"testing"
)

type xorOperationTestPair struct {
	inputN     int
	inputStart int
	out        int
}

var xorOperationTestCases = []xorOperationTestPair{
	{5, 0, 8},
	{4, 3, 8},
}

func TestEvalXorOperation(t *testing.T) {
	for _, pair := range xorOperationTestCases {
		var newOut = xorOperation(pair.inputN, pair.inputStart)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputN,
				"and", pair.inputStart,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
