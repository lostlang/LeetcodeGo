package anumberafteradoublereversal

import (
	"reflect"
	"testing"
)

type isSameAfterReversalsTestPair struct {
	input  int
	output bool
}

var isSameAfterReversalsTestCases = []isSameAfterReversalsTestPair{
	{526, true},
	{1800, false},
	{0, true},
}

func TestEvalIsSameAfterReversals(t *testing.T) {
	for _, pair := range isSameAfterReversalsTestCases {
		newOut := isSameAfterReversals(pair.input)
		if !reflect.DeepEqual(newOut, pair.output) {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", newOut,
			)
		}
	}
}
