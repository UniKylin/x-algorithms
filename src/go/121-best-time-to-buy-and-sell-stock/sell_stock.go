// link		: https://leetcode-cn.com/problems/subsets/
// Author	: Kylin
// Date		: 2021-10-13

package leetcode

func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([]int, 2)
	dp[0] = 0
	dp[1] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[0] = max(dp[0], dp[1]+prices[i])
		dp[1] = max(dp[1], -prices[i])
	}
	return dp[0]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
