package BinaryNumberWithAlternatingBits

import (
	"reflect"
	"testing"
)

type hasAlternatingBitsTestPair struct {
	input  int
	output bool
}

var hasAlternatingBitsTestCases = []hasAlternatingBitsTestPair{
	{5, true},
	{7, false},
	{11, false},
}

func TestEvalHasAlternatingBits(t *testing.T) {
	for _, pair := range hasAlternatingBitsTestCases {
		newOutput := hasAlternatingBits(pair.input)
		if !reflect.DeepEqual(newOutput, pair.output) {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", newOutput,
			)
		}
	}
}
