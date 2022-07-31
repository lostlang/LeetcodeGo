package leetcode

import (
	"reflect"
	"testing"
)

type freqAlphabetsTestPair struct {
	input string
	out   string
}

var freqAlphabetsTestCases = []freqAlphabetsTestPair{
	{"10#11#12", "jkab"},
	{"1326#", "acz"},
}

func TestEvalFreqAlphabets(t *testing.T) {
	for _, pair := range freqAlphabetsTestCases {
		var newOut = freqAlphabets(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
