package leetcode

func canJump(nums []int) bool {
	var mapJump = make(map[int]bool)
	mapJump[0] = true

	for i, v := range nums {
		if mapJump[i] {
			for j := 1; j <= v; j++ {
				mapJump[i+j] = true
			}
			if mapJump[len(nums)-1] {
				return true
			}
		}
	}

	return false
}
