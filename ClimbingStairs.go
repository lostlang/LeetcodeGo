package leetcode

func climbStairs(n int) int {
	var hashC = make(map[int]int)

	return hashClimb(n, hashC)
}

func hashClimb(num int, hash map[int]int) int {
	if hash[num] != 0 {
		return hash[num]
	}
	if num == 0 {
		return 1
	}

	if num < 0 {
		return 0
	}

	var left = hashClimb(num-1, hash)
	var right = hashClimb(num-2, hash)

	hash[num] = left + right

	return hash[num]
}
