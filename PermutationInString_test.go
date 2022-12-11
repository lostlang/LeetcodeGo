package leetcode

import (
	"reflect"
	"testing"
)

type checkInclusionTestPair struct {
	inputS1 string
	inputS2 string
	out     bool
}

var checkInclusionTestCases = []checkInclusionTestPair{
	{"ab", "eidbaooo", true},
	{"ab", "eidboaoo", false},
	{"a", "a", true},
	{"a", "b", false},
}

func TestEvalCheckInclusion(t *testing.T) {
	for _, pair := range checkInclusionTestCases {
		var newOut = checkInclusion(pair.inputS1, pair.inputS2)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputS1, pair.inputS2,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
