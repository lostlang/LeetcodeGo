package leetcode

import (
	"reflect"
	"testing"
)

type repeatedCharacterTestPair struct {
	input string
	out   byte
}

var repeatedCharacterTestCases = []repeatedCharacterTestPair{
	{"abccbaacz", 'c'},
	{"abcdd", 'd'},
}

func TestEvalRepeatedCharacter(t *testing.T) {
	for _, pair := range repeatedCharacterTestCases {
		var newOut = repeatedCharacter(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
