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
		newOutput := addBinary(pair.inputA, pair.inputB)
		if !reflect.DeepEqual(newOutput, pair.output) {
			t.Error(
				"For", pair.inputA, pair.inputB,
				"expected", pair.output,
				"got", newOutput,
			)
		}
	}
}
