package CheckIfNAndItsDoubleExist

import (
	"reflect"
	"testing"
)

type checkIfExistTestPair struct {
	input  []int
	output bool
}

var checkIfExistTestCases = []checkIfExistTestPair{
	{[]int{10, 2, 5, 3}, true},
	{[]int{7, 1, 14, 11}, true},
	{[]int{3, 1, 7, 11}, false},
}

func TestEvalCheckIfExist(t *testing.T) {
	for _, pair := range checkIfExistTestCases {
		newOut := checkIfExist(pair.input)
		if !reflect.DeepEqual(newOut, pair.output) {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", newOut,
			)
		}
	}
}
