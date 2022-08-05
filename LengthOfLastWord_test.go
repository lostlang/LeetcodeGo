package leetcode

import (
	"reflect"
	"testing"
)

type lengthOfLastWordTestPair struct {
	input string
	out   int
}

var lengthOfLastWordTestCases = []lengthOfLastWordTestPair{
	{"Hello World", 5},
	{"   fly me   to   the moon  ", 4},
	{"luffy is still joyboy", 6},
	{"Today is a nice day", 3},
}

func TestEvalLengthOfLastWord(t *testing.T) {
	for _, pair := range lengthOfLastWordTestCases {
		var newOut = lengthOfLastWord(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
