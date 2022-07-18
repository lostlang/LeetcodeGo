package leetcode

// 724. Find Pivot Index

func findMiddleIndex(nums []int) int {
	var sum, sum2, tmp int

	for _, val := range nums {
		sum += val
	}

	for i, val := range nums {
		sum2 += tmp
		sum -= val

		if sum == sum2 {
			return i
		}

		tmp = val
	}

	return -1
}
