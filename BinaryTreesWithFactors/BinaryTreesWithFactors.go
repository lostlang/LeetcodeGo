package BinaryTreesWithFactors

import (
	"sort"
)

func numFactoredBinaryTrees(arr []int) int {
	var out int
	hash := make(map[int]int)
	sort.Ints(arr)
	mod := 1_000_000_007

	for i, v := range arr {
		for j := 0; j < i; j++ {
			if v%arr[j] == 0 {
				hash[v] += hash[arr[j]] * hash[v/arr[j]]
				hash[v] %= mod
			}
		}

		hash[v] += 1
		hash[v] %= mod
	}

	for _, v := range hash {
		out += v
		out %= mod
	}

	return out
}
