package leetcode

func getMaxLen(nums []int) int {
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
			if v[0] <= 0 {
				maxForLine = append(maxForLine, 0)
			} else {
				maxForLine = append(maxForLine, 1)
			}
			continue
		}

		var negatives = countNegativesInLine(v)
		if negatives%2 != 0 {
			var start, end = searchNegatives(v)
			maxForLine = append(maxForLine,
				maxInLine([]int{len(v[:end]), len(v[start+1:])}))
		} else {
			maxForLine = append(maxForLine, len(v))
		}
	}

	out = maxInLine(maxForLine)

	return out
}
