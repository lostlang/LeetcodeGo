package leetcode

func lengthOfLIS(nums []int) int {
	var maxDist int
	distance := make([]int, len(nums))
	distance[0] = 1
	out := 1

	for i := 1; i < len(nums); i++ {
		maxDist = 0

		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && distance[j] > maxDist {
				maxDist = distance[j]
			}
		}

		distance[i] = maxDist + 1
		if distance[i] > out {
			out = distance[i]
		}
	}

	return out
}
