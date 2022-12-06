package leetcode

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	var out [][]string
	var mapStrs = make(map[string][]string)

	for _, str := range strs {
		var sortedStr = sortString(str)
		mapStrs[sortedStr] = append(mapStrs[sortedStr], str)
	}

	for _, v := range mapStrs {
		out = append(out, v)
	}

	return out
}

func sortString(s string) string {
	var a = []rune(s)

	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	return string(a)
}
