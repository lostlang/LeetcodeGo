package addtoarrayformofinteger

import (
	"reflect"
	"testing"
)

type addToArrayFormTestPair struct {
	inputArr []int
	inputK   int
	output   []int
}

var addToArrayFormTestCases = []addToArrayFormTestPair{
	{[]int{1, 2, 0, 0}, 34, []int{1, 2, 3, 4}},
	{[]int{2, 7, 4}, 181, []int{4, 5, 5}},
	{[]int{2, 1, 5}, 806, []int{1, 0, 2, 1}},
}

func TestEvalAddToArrayForm(t *testing.T) {
	for _, pair := range addToArrayFormTestCases {
		newOut := addToArrayForm(pair.inputArr, pair.inputK)
		if !reflect.DeepEqual(newOut, pair.output) {
			t.Error(
				"For", pair.inputArr,
				"and", pair.inputK,
				"expected", pair.output,
				"got", newOut,
			)
		}
	}
}
