package leetcode

import (
	"reflect"
	"testing"
)

type evalRPNTestPair struct {
	input []string
	out   int
}

var evalRPNTestCases = []evalRPNTestPair{
	{[]string{"2", "1", "+", "3", "*"}, 9},
	{[]string{"4", "13", "5", "/", "+"}, 6},
	{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
}

func TestEvalEvalRPN(t *testing.T) {
	for _, pair := range evalRPNTestCases {
		var newOut = evalRPN(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
