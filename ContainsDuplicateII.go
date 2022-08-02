package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {
	var hash = make(map[int]int)

	for i, val := range nums {
		if hash[val] == 0 {
			var tmp = len(hash)
			hash[val] = 0

			if len(hash) == tmp {
				if i-hash[val] <= k {
					return true
				}
			}
		} else {
			if i-hash[val] <= k {
				return true
			}
		}

		hash[val] = i
	}

	return false
}
