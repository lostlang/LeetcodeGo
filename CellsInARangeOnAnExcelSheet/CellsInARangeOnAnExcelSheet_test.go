package CellsInARangeOnAnExcelSheet

import (
	"reflect"
	"testing"
)

type cellsInRangeTestPair struct {
	input  string
	output []string
}

var cellsInRangeTestCases = []cellsInRangeTestPair{
	{"K1:L2", []string{"K1", "K2", "L1", "L2"}},
	{"A1:F1", []string{"A1", "B1", "C1", "D1", "E1", "F1"}},
}

func TestEvalCellsInRange(t *testing.T) {
	for _, pair := range cellsInRangeTestCases {
		newOut := cellsInRange(pair.input)
		if !reflect.DeepEqual(newOut, pair.output) {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", newOut,
			)
		}
	}
}
