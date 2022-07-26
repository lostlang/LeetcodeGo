package leetcode

func maximum69Number(num int) int {
	var out, tmp int
	var chance = true

	for num > 0 {
		out = out*10 + num%10
		num /= 10
	}

	for out > 0 {
		tmp = out % 10
		if chance && tmp == 6 {
			chance = false
			tmp = 9
		}

		num = num*10 + tmp
		out /= 10
	}

	return num
}
