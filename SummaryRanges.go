package leetcode

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	var out []string
	var start = nums[0]
	var end = nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] == end {
			continue
		}

		if nums[i] == end+1 {
			end = nums[i]
			continue
		}

		out = append(out, newRange(start, end))

		start = nums[i]
		end = nums[i]
	}

	out = append(out, newRange(start, end))

	return out
}

func newRange(start, end int) string {
	if start == end {
		return fmt.Sprintf("%d", start)
	}

	return fmt.Sprintf("%d->%d", start, end)
}
