package leetcode

import (
	"reflect"
	"testing"
)

type checkIfPangramTestPair struct {
	input string
	out   bool
}

var checkIfPangramTestCases = []checkIfPangramTestPair{
	{"thequickbrownfoxjumpsoverthelazydog", true},
	{"leetcode", false},
}

func TestEvalCheckIfPangram(t *testing.T) {
	for _, pair := range checkIfPangramTestCases {
		newOut := checkIfPangram(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
