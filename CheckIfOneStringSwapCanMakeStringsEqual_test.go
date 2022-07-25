package leetcode

import (
	"reflect"
	"testing"
)

type areAlmostEqualTestPair struct {
	inputS1 string
	inputS2 string
	out     bool
}

var areAlmostEqualTestCases = []areAlmostEqualTestPair{
	{"bank", "kanb", true},
	{"attack", "defend", false},
	{"kelb", "kelb", true},
}

func TestEvalAreAlmostEqual(t *testing.T) {
	for _, pair := range areAlmostEqualTestCases {
		var newOut = areAlmostEqual(pair.inputS1, pair.inputS2)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputS1,
				"and", pair.inputS2,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
