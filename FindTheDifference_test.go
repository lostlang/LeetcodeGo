package leetcode

import (
	"reflect"
	"testing"
)

type findTheDifferenceTestPair struct {
	inputS string
	inputT string
	out    byte
}

var findTheDifferenceTestCases = []findTheDifferenceTestPair{
	{"abcd", "abcde", 'e'},
	{"", "e", 'e'},
}

func TestEvalFindTheDifference(t *testing.T) {
	for _, pair := range findTheDifferenceTestCases {
		var newOut = findTheDifference(pair.inputS, pair.inputT)
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
