package adddigits

import (
	"reflect"
	"testing"
)

type addDigitsTestPair struct {
	input  int
	output int
}

var addDigitsTestCases = []addDigitsTestPair{
	{38, 2},
	{0, 0},
}

func TestEvalAddDigits(t *testing.T) {
	for _, pair := range addDigitsTestCases {
		newOut := addDigits(pair.input)
		if !reflect.DeepEqual(newOut, pair.output) {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", newOut,
			)
		}
	}
}