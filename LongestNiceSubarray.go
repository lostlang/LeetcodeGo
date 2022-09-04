package leetcode

func longestNiceSubarray(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	var out int
	var tmpArr []int

	for _, num := range nums {
		if simpleAll(tmpArr, num) {
			tmpArr = append(tmpArr, num)
		} else {
			if len(tmpArr) > out {
				out = len(tmpArr)
			}

			for !simpleAll(tmpArr[1:], num) {
				tmpArr = tmpArr[1:]
			}

			tmpArr = append(tmpArr, num)
		}
	}

	if len(tmpArr) > out {
		out = len(tmpArr)
	}

	return out - 1
}

func simpleAll(nums []int, num int) bool {
	if len(nums) == 0 {
		return true
	}

	for _, num2 := range nums {
		if num2&num != 0 {
			return false
		}
	}

	return true
}
