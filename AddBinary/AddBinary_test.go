package addbinary

import (
	"reflect"
	"testing"
)

type addBinaryTestPair struct {
	inputA string
	inputB string
	output string
}

var addBinaryTestCases = []addBinaryTestPair{
	{"11", "1", "100"},
	{"1010", "1011", "10101"},
	{"0", "0", "0"},
}

func TestEvalAddBinary(t *testing.T) {
	for _, pair := range addBinaryTestCases {
		newOut := addBinary(pair.inputA, pair.inputB)
		if !reflect.DeepEqual(newOut, pair.output) {
			t.Error(
				"For", pair.inputA, pair.inputB,
				"expected", pair.output,
				"got", newOut,
			)
		}
	}
}
