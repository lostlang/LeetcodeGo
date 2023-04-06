package AverageSalaryExcludingTheMinimumAndMaximumSalary

func average(salary []int) float64 {
	maxNum, minNum := salary[0], salary[0]
	sumArr := 0

	for _, val := range salary {
		sumArr += val

		if val > maxNum {
			maxNum = val
		}

		if val < minNum {
			minNum = val
		}
	}

	return float64(sumArr-minNum-maxNum) / float64(len(salary)-2)
}
