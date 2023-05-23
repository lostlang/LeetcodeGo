package BuddyStrings

import (
	"reflect"
	"testing"
)

type buddyStringsTestPair struct {
	inputS    string
	inputGoal string
	out       bool
}

var buddyStringsTestCases = []buddyStringsTestPair{
	{"ab", "ba", true},
	{"ab", "ab", false},
	{"aa", "aa", true},
}

func TestEvalBuddyStrings(t *testing.T) {
	for _, pair := range buddyStringsTestCases {
		newOut := buddyStrings(pair.inputS, pair.inputGoal)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputS,
				"and", pair.inputGoal,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
