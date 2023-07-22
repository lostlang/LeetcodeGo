package BestTimeToBuyAndSellStock

func maxProfit(prices []int) int {
	output := 0
	lowPrice := prices[0]

	for _, val := range prices {
		if val < lowPrice {
			lowPrice = val
		}

		if output < val-lowPrice {
			output = val - lowPrice
		}
	}

	return output
}
