package leetcode

import (
	"reflect"
	"testing"
)

type reverseStringTestPair struct {
	input []byte
	out   []byte
}

var reverseStringTestCases = []reverseStringTestPair{
	{[]byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
	{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}, []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
}

func TestEvalReverseString(t *testing.T) {
	for _, pair := range reverseStringTestCases {
		reverseString(pair.input)
		if !reflect.DeepEqual(pair.input, pair.out) {
			t.Error(
				"Expected", pair.out,
				"got", pair.input,
			)
		}
	}
}
