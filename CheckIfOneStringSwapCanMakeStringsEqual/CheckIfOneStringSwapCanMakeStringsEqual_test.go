package CheckIfOneStringSwapCanMakeStringsEqual

import (
	"reflect"
	"testing"
)

type areAlmostEqualTestPair struct {
	inputS1 string
	inputS2 string
	output  bool
}

var areAlmostEqualTestCases = []areAlmostEqualTestPair{
	{"bank", "kanb", true},
	{"attack", "defend", false},
	{"kelb", "kelb", true},
}

func TestEvalAreAlmostEqual(t *testing.T) {
	for _, pair := range areAlmostEqualTestCases {
		newOut := areAlmostEqual(pair.inputS1, pair.inputS2)
		if !reflect.DeepEqual(newOut, pair.output) {
			t.Error(
				"For", pair.inputS1, pair.inputS2,
				"expected", pair.output,
				"got", newOut,
			)
		}
	}
}
