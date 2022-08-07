package leetcode

import (
	"reflect"
	"testing"
)

type decodeMessageTestPair struct {
	inputKey     string
	inputMessage string
	out          string
}

var decodeMessageTestCases = []decodeMessageTestPair{
	{"the quick brown fox jumps over the lazy dog", "vkbs bs t suepuv", "this is a secret"},
	{"eljuxhpwnyrdgtqkviszcfmabo", "zwx hnfx lqantp mnoeius ycgk vcnjrdb", "the five boxing wizards jump quickly"},
}

func TestEvalDecodeMessage(t *testing.T) {
	for _, pair := range decodeMessageTestCases {
		var newOut = decodeMessage(pair.inputKey, pair.inputMessage)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputKey,
				"and", pair.inputMessage,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
