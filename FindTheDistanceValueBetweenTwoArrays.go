package leetcode

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	var out = len(arr1)

	for _, val := range arr1 {
		for _, val2 := range arr2 {
			if val-val2 <= d && val2-val <= d {
				out--
				break
			}
		}
	}

	return out
}
