package leetcode

import (
	"reflect"
	"testing"
)

type truncateSentenceTestPair struct {
	inputS string
	inputK int
	out    string
}

var truncateSentenceTestCases = []truncateSentenceTestPair{
	{"Hello how are you Contestant", 4, "Hello how are you"},
	{"What is the solution to this problem", 4, "What is the solution"},
	{"chopper is not a tanuki", 5, "chopper is not a tanuki"},
}

func TestEvalTruncateSentence(t *testing.T) {
	for _, pair := range truncateSentenceTestCases {
		newOut := truncateSentence(pair.inputS, pair.inputK)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputS, pair.inputK,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
