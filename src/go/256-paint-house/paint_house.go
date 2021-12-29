// link		: https://leetcode-cn.com/problems/paint-house/
// Author	: Kylin
// Date		: 2021-12-29

package leetcode

func MinCost(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}
	for i := 1; i < len(costs); i++ {
		costs[i][0] = min(costs[i-1][1], costs[i-1][2]) + costs[i][0]
		costs[i][1] = min(costs[i-1][0], costs[i-1][2]) + costs[i][1]
		costs[i][2] = min(costs[i-1][0], costs[i-1][1]) + costs[i][2]
	}
	return min(costs[len(costs)-1]...)
}

func min(a ...int) int {
	res := a[0]
	for k := range a {
		if a[k] < res {
			res = a[k]
		}
	}
	return res
}
