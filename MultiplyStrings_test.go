package leetcode

import (
	"reflect"
	"testing"
)

type multiplyTestPair struct {
	inputS1 string
	inputS2 string
	out     string
}

var multiplyTestCases = []multiplyTestPair{
	// {"2", "3", "6"},
	// {"123", "456", "56088"},
}

func TestEvalMultiply(t *testing.T) {
	for _, pair := range multiplyTestCases {
		var newOut = multiply(pair.inputS1, pair.inputS2)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputS1,
				"and", pair.inputS2,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
