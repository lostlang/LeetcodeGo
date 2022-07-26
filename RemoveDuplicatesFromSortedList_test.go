package leetcode

import (
	"reflect"
	"testing"
)

type deleteDuplicatesTestPair struct {
	input *ListNode
	out   *ListNode
}

var deleteDuplicatesTestCases = []deleteDuplicatesTestPair{
	{generateListNode([]int{1, 1, 2}), generateListNode([]int{1, 2})},
	{generateListNode([]int{1, 1, 2, 3, 3}), generateListNode([]int{1, 2, 3})},
}

func TestEvalDeleteDuplicates(t *testing.T) {
	for _, pair := range deleteDuplicatesTestCases {
		var newOut = deleteDuplicates(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
