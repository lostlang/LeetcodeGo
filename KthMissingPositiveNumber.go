package leetcode

func findKthPositive(arr []int, k int) int {
	if len(arr)+k > arr[len(arr)-1] {
		return len(arr) + k
	}

	if arr[0] > k {
		return k
	}

	k -= arr[0] - 1

	for i := 1; i < len(arr); i++ {
		if arr[i-1]+k < arr[i] {
			return arr[i-1] + k
		}

		k -= arr[i] - arr[i-1] - 1
	}

	return 0
}
