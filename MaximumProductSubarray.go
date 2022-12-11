package leetcode

func maxProduct(nums []int) int {
	var out int
	var numsSplit = make([][]int, 0)
	var line = make([]int, 0)
	for _, v := range nums {
		if v == 0 {
			if len(line) > 0 {
				numsSplit = append(numsSplit, line)
			}
			numsSplit = append(numsSplit, []int{0})
			line = make([]int, 0)
		} else {
			line = append(line, v)
		}
	}
	if len(line) > 0 {
		numsSplit = append(numsSplit, line)
	}

	var maxForLine = make([]int, 0)

	for _, v := range numsSplit {
		if len(v) == 1 {
			maxForLine = append(maxForLine, v[0])
			continue
		}

		var negatives = countNegativesInLine(v)
		if negatives%2 != 0 {
			var start, end = searchNegatives(v)
			maxForLine = append(maxForLine,
				maxInLine([]int{produceLine(v[:end]),
					produceLine(v[start+1:])}))
		} else {
			maxForLine = append(maxForLine, produceLine(v))
		}
	}

	out = maxInLine(maxForLine)

	return out
}

func countNegativesInLine(nums []int) int {
	var out = 0
	for _, v := range nums {
		if v < 0 {
			out++
		}
	}

	return out
}

func searchNegatives(nums []int) (int, int) {
	var start, end = -1, -1
	for i, v := range nums {
		if v < 0 {
			if start == -1 {
				start = i
			}
			end = i
		}
	}

	return start, end
}

func produceLine(nums []int) int {
	var out = 1
	for _, v := range nums {
		out *= v
	}

	return out
}

func maxInLine(nums []int) int {
	var out = nums[0]
	for _, v := range nums {
		if v > out {
			out = v
		}
	}

	return out
}
