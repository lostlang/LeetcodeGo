package leetcode

import (
	"reflect"
	"testing"
)

type sumPrefixScoresTestPair struct {
	input []string
	out   []int
}

var sumPrefixScoresTestCases = []sumPrefixScoresTestPair{
	{[]string{"abc", "ab", "bc", "b"}, []int{5, 4, 3, 2}},
	{[]string{"abcd"}, []int{4}},
	{[]string{"qtcqcmwcin", "vkjotbrbzn", "eoorlyfche", "eoorlyhn", "eoorlyfcxk", "qfnmjilcom", "eoorlyfche", "qtcqcmwcnl", "qtcqcrpjr"}, []int{24, 10, 34, 26, 32, 13, 34, 24, 20}},
}

func TestEvalSumPrefixScores(t *testing.T) {
	for _, pair := range sumPrefixScoresTestCases {
		var newOut = sumPrefixScores(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
