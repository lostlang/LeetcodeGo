package CheckIfTheSentenceIsPangram

import (
	"reflect"
	"testing"
)

type checkIfPangramTestPair struct {
	input  string
	output bool
}

var checkIfPangramTestCases = []checkIfPangramTestPair{
	{"thequickbrownfoxjumpsoverthelazydog", true},
	{"leetcode", false},
}

func TestEvalCheckIfPangram(t *testing.T) {
	for _, pair := range checkIfPangramTestCases {
		newOut := checkIfPangram(pair.input)
		if !reflect.DeepEqual(newOut, pair.output) {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", newOut,
			)
		}
	}
}
