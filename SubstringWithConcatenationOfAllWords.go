package leetcode

import (
	"sort"
	"strings"
)

func findSubstring(s string, words []string) []int {
	var index []map[int]bool
	var out []int

	for _, word := range words {
		index = append(index, findAll(s, word))
	}

	for start := range index[0] {
		out = append(out, findPairs(start, start, len(words[0]), index[1:])...)
	}

	hash := make(map[int]bool)

	for _, i := range out {
		hash[i] = true
	}

	out = make([]int, len(hash))
	var j int

	for i := range hash {
		out[j] = i
		j++
	}

	sort.Ints(out)

	return out
}

func findAll(s string, word string) map[int]bool {
	var out = make(map[int]bool)

	for i := 0; i < len(s); i++ {
		if newOut := strings.Index(s[i:], word); newOut != -1 {
			i += newOut
			out[i] = true
		}
	}

	return out
}

func findPairs(min, max, lenWord int, words []map[int]bool) []int {
	if len(words) == 0 {
		return []int{min}
	}

	var out []int

	for i, word := range words {
		var tmp = make([]map[int]bool, len(words)-1)
		copy(tmp, words[:i])
		copy(tmp[i:], words[i+1:])

		if word[min-lenWord] {
			out = append(out, findPairs(min-lenWord, max, lenWord, tmp)...)
		}

		if word[max+lenWord] {
			out = append(out, findPairs(min, max+lenWord, lenWord, tmp)...)
		}

	}

	return out
}
