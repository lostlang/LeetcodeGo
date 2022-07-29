package leetcode

import (
	"reflect"
	"testing"
)

type interpretTestPair struct {
	input string
	out   string
}

var interpretTestCases = []interpretTestPair{
	{"G()(al)", "Goal"},
	{"G()()()()(al)", "Gooooal"},
	{"(al)G(al)()()G", "alGalooG"},
}

func TestEvalInterpret(t *testing.T) {
	for _, pair := range interpretTestCases {
		var newOut = interpret(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
