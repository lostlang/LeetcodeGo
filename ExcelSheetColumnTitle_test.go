package leetcode

import (
	"reflect"
	"testing"
)

type convertToTitleTestPair struct {
	input int
	out   string
}

var convertToTitleTestCases = []convertToTitleTestPair{
	{1, "A"},
	{28, "AB"},
	{701, "ZY"},
}

func TestEvalConvertToTitle(t *testing.T) {
	for _, pair := range convertToTitleTestCases {
		var newOut = convertToTitle(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
