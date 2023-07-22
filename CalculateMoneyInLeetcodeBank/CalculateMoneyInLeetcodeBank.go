package CalculateMoneyInLeetcodeBank

func totalMoney(n int) int {
	pair := n / 7
	unpair := n % 7
	pairSum := pair*28 + (pair-1)*pair/2*7
	unpairSum := (1+unpair)*unpair/2 + pair*unpair

	return pairSum + unpairSum
}
