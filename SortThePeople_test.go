package leetcode

import (
	"reflect"
	"testing"
)

type sortPeopleTestPair struct {
	inputNames   []string
	inputHeights []int
	out          []string
}

var sortPeopleTestCases = []sortPeopleTestPair{
	{[]string{"Mary", "John", "Emma"}, []int{180, 165, 170}, []string{"Mary", "Emma", "John"}},
	{[]string{"Alice", "Bob", "Bob"}, []int{155, 185, 150}, []string{"Bob", "Alice", "Bob"}},
}

func TestEvalSortPeople(t *testing.T) {
	for _, pair := range sortPeopleTestCases {
		newOut := sortPeople(pair.inputNames, pair.inputHeights)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputNames, pair.inputHeights,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
