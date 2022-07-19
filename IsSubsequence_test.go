package leetcode

import (
	"reflect"
	"testing"
)

type isSubsequenceTestPair struct {
	inputS string
	inputT string
	out    bool
}

var isSubsequenceTestCases = []isSubsequenceTestPair{
	{"abc", "ahbgdc", true},
	{"axc", "ahbgdc", false},
	{"aaaaaa", "bbaaaa", false},
}

func TestEvalIsSubsequence(t *testing.T) {
	for _, pair := range isSubsequenceTestCases {
		var newOut = isSubsequence(pair.inputS, pair.inputT)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputS,
				"and", pair.inputT,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
