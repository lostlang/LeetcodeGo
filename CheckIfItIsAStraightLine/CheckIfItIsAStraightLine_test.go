package CheckIfItIsAStraightLine

import (
	"reflect"
	"testing"
)

type checkStraightLineTestPair struct {
	input  [][]int
	output bool
}

var checkStraightLineTestCases = []checkStraightLineTestPair{
	{[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}, true},
	{[][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}, false},
}

func TestEvalCheckStraightLine(t *testing.T) {
	for _, pair := range checkStraightLineTestCases {
		newOutput := checkStraightLine(pair.input)
		if !reflect.DeepEqual(newOutput, pair.output) {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", newOutput,
			)
		}
	}
}
