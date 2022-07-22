package leetcode

func maxProfit(prices []int) int {
	var out int
	var lowPrice = prices[0]

	for _, val := range prices {
		if val < lowPrice {
			lowPrice = val
		}

		if out < val-lowPrice {
			out = val - lowPrice
		}
	}

	return out
}
