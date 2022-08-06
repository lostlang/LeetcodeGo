package leetcode

import (
	"reflect"
	"testing"
)

type countMatchesTestPair struct {
	inputItems     [][]string
	inputRuleKey   string
	inputRuleValue string
	out            int
}

var countMatchesTestCases = []countMatchesTestPair{
	{[][]string{[]string{"phone", "blue", "pixel"}, []string{"computer", "silver", "lenovo"}, []string{"phone", "gold", "iphone"}}, "color", "silver", 1},
	{[][]string{[]string{"phone", "blue", "pixel"}, []string{"computer", "silver", "lenovo"}, []string{"phone", "gold", "iphone"}}, "type", "phone", 2},
}

func TestEvalCountMatches(t *testing.T) {
	for _, pair := range countMatchesTestCases {
		var newOut = countMatches(pair.inputItems, pair.inputRuleKey, pair.inputRuleValue)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputItems,
				",", pair.inputRuleKey,
				"and", pair.inputRuleValue,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
