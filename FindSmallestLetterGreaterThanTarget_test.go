package leetcode

import (
	"reflect"
	"testing"
)

type nextGreatestLetterTestPair struct {
	inputLetters []byte
	inputTarget  byte
	out          byte
}

var nextGreatestLetterTestCases = []nextGreatestLetterTestPair{
	{[]byte{'c', 'f', 'j'}, 'a', 'c'},
	{[]byte{'c', 'f', 'j'}, 'c', 'f'},
	{[]byte{'c', 'f', 'j'}, 'd', 'f'},
	{[]byte{'a', 'b'}, 'z', 'a'},
}

func TestEvalNextGreatestLetter(t *testing.T) {
	for _, pair := range nextGreatestLetterTestCases {
		var newOut = nextGreatestLetter(pair.inputLetters, pair.inputTarget)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputLetters,
				"and", pair.inputTarget,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
