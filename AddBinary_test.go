package leetcode

import (
	"reflect"
	"testing"
)

type addBinaryTestPair struct {
	inputA string
	inputB string
	out    string
}

var addBinaryTestCases = []addBinaryTestPair{
	{"11", "1", "100"},
	{"1010", "1011", "10101"},
	{"0", "0", "0"},
}

func TestEvalAddBinary(t *testing.T) {
	for _, pair := range addBinaryTestCases {
		var newOut = addBinary(pair.inputA, pair.inputB)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputA,
				"and", pair.inputB,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
