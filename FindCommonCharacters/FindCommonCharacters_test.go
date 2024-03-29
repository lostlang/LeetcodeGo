package FindCommonCharacters

import (
	"reflect"
	"sort"
	"testing"
)

type commonCharsTestPair struct {
	input []string
	out   []string
}

var commonCharsTestCases = []commonCharsTestPair{
	{[]string{"bella", "label", "roller"}, []string{"e", "l", "l"}},
	{[]string{"cool", "lock", "cook"}, []string{"c", "o"}},
}

func TestEvalCommonChars(t *testing.T) {
	for _, pair := range commonCharsTestCases {
		newOut := commonChars(pair.input)
		sort.Strings(newOut)
		sort.Strings(pair.out)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
