package leetcode

func distributeCandies(candies int, num_people int) []int {
	var out = make([]int, num_people)
	var index int
	var candy = 1

	for candies > 0 {
		if candy > candies {
			candy = candies
		}
		out[index] += candy
		candies -= candy
		candy++
		index++
		if index == num_people {
			index = 0
		}
	}
	return out
}
