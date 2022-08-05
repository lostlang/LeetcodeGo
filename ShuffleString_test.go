package leetcode

import (
	"reflect"
	"testing"
)

type restoreStringTestPair struct {
	inputS string
	inputI []int
	out    string
}

var restoreStringTestCases = []restoreStringTestPair{
	{"codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3}, "leetcode"},
	{"abc", []int{0, 1, 2}, "abc"},
}

func TestEvalRestoreString(t *testing.T) {
	for _, pair := range restoreStringTestCases {
		var newOut = restoreString(pair.inputS, pair.inputI)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputS,
				"and", pair.inputI,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
