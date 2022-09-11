package leetcode

import (
	"sort"
)

func minGroups(intervals [][]int) int {
	var out int

	var hash = make(map[int]int)

	for _, interval := range intervals {
		hash[interval[0]]++
		hash[interval[1]+1]--
	}

	var keys = make([]int, 0, len(hash))
	for k := range hash {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for i := range keys {
		if i > 0 {
			hash[keys[i]] += hash[keys[i-1]]
		}
	}

	for _, v := range hash {
		if v > out {
			out = v
		}
	}

	return out
}
