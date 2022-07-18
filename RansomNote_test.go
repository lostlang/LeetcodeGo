package leetcode

import (
	"reflect"
	"testing"
)

type canConstructTestPair struct {
	inputNote     string
	inputMagazine string
	out           bool
}

var canConstructTestCases = []canConstructTestPair{
	{"a", "b", false},
	{"aa", "aab", true},
}

func TestEvalCanConstruct(t *testing.T) {
	for _, pair := range canConstructTestCases {
		var newOut = canConstruct(pair.inputNote, pair.inputMagazine)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputNote,
				"and", pair.inputMagazine,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
