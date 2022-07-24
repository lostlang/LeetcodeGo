package leetcode

import (
	"reflect"
	"testing"
)

type isAnagramTestPair struct {
	inputS string
	inputT string
	out    bool
}

var isAnagramTestCases = []isAnagramTestPair{
	{"anagram", "nagaram", true},
	{"rat", "car", false},
}

func TestEvalIsAnagram(t *testing.T) {
	for _, pair := range isAnagramTestCases {
		var newOut = isAnagram(pair.inputS, pair.inputT)
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
