package leetcode

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var out = make([]bool, len(candies))

	var maxCandie = candies[0]

	for _, candie := range candies {
		if candie > maxCandie {
			maxCandie = candie
		}
	}

	for index := range candies {
		if candies[index]+extraCandies >= maxCandie {
			out[index] = true
		}
	}

	return out
}
