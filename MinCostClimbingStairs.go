package leetcode

func minCostClimbingStairs(cost []int) int {
	for i := 2; i < len(cost); i++ {
		if cost[i-1] > cost[i-2] {
			cost[i] += cost[i-2]
		} else {
			cost[i] += cost[i-1]
		}
	}

	if cost[len(cost)-1] < cost[len(cost)-2] {
		return cost[len(cost)-1]
	}

	return cost[len(cost)-2]
}
