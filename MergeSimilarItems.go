package leetcode

import (
	"sort"
)

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	var hash = make(map[int]int)

	for _, val := range items1 {
		hash[val[0]] += val[1]
	}

	for _, val := range items2 {
		hash[val[0]] += val[1]
	}

	var keys = make([]int, len(hash))
	var out = make([][]int, len(hash))
	var i int

	for val := range hash {
		keys[i] = val
		i++
	}

	sort.Ints(keys)
	i = 0

	for val := range keys {
		out[i] = []int{keys[val], hash[keys[val]]}
		i++
	}

	return out
}
