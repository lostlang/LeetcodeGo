package leetcode

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	var pigs int
	var tests = minutesToTest/minutesToDie + 1
	var tmp = 1

	for tmp < buckets {
		pigs++
		tmp *= tests
	}

	return pigs
}
