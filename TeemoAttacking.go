package leetcode

func findPoisonedDuration(timeSeries []int, duration int) int {
	var out int

	for i := 1; i < len(timeSeries); i++ {
		if timeSeries[i]-timeSeries[i-1] > duration {
			out += duration
		} else {
			out += timeSeries[i] - timeSeries[i-1]
		}
	}

	out += duration
	return out
}
