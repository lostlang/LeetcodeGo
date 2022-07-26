package leetcode

func totalMoney(n int) int {
	var pair = n / 7
	var unpair = n % 7
	var pairSum = pair*28 + (pair-1)*pair/2*7
	var unpairSum = (1+unpair)*unpair/2 + pair*unpair

	return pairSum + unpairSum
}
