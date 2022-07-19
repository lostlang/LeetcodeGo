package leetcode

import (
	"reflect"
	"testing"
)

type isIsomorphicTestPair struct {
	inputS string
	inputT string
	out    bool
}

var isIsomorphicTestCases = []isIsomorphicTestPair{
	{"egg", "add", true},
	{"foo", "bar", false},
	{"paper", "title", true},
}

func TestEvalIsIsomorphic(t *testing.T) {
	for _, pair := range isIsomorphicTestCases {
		var newOut = isIsomorphic(pair.inputS, pair.inputT)
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
