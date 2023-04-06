package addtwointegers

import (
	"reflect"
	"testing"
)

type sumTestPair struct {
	inputA int
	inputB int
	output int
}

var sumTestCases = []sumTestPair{
	{12, 5, 17},
	{-10, 4, -6},
}

func TestEvalsum(t *testing.T) {
	for _, pair := range sumTestCases {
		newOutput := sum(pair.inputA, pair.inputB)
		if !reflect.DeepEqual(newOutput, pair.output) {
			t.Error(
				"For", pair.inputA, pair.inputB,
				"expected", pair.output,
				"got", newOutput,
			)
		}
	}
}
