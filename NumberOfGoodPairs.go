package leetcode

func numIdenticalPairs(nums []int) int {
	var out int

	var numHash = make(map[int]int)

	for _, num := range nums {
		numHash[num]++
	}

	for num := range numHash {
		out += numHash[num] * (numHash[num] - 1) / 2
	}
	return out
}
