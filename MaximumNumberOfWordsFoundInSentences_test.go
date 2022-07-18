package leetcode

import (
	"reflect"
	"testing"
)

type mostWordsFoundTestPair struct {
	input []string
	out   int
}

var mostWordsFoundTestCases = []mostWordsFoundTestPair{
	{[]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}, 6},
	{[]string{"please wait", "continue to fight", "continue to win"}, 3},
}

func TestEvalMostWordsFound(t *testing.T) {
	for _, pair := range mostWordsFoundTestCases {
		var newOut = mostWordsFound(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
