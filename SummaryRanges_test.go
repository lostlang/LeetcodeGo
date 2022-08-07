package leetcode

import (
	"reflect"
	"testing"
)

type summaryRangesTestPair struct {
	input []int
	out   []string
}

var summaryRangesTestCases = []summaryRangesTestPair{
	{[]int{0, 1, 2, 4, 5, 7}, []string{"0->2", "4->5", "7"}},
	{[]int{0, 2, 3, 4, 6, 8, 9}, []string{"0", "2->4", "6", "8->9"}},
}

func TestEvalSummaryRanges(t *testing.T) {
	for _, pair := range summaryRangesTestCases {
		var newOut = summaryRanges(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
