// link		: https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
// Author	: Kylin
// Date		: 2021-12-06

package leetcode

func MaxProfitCooldown(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	// 三个状态，天数，是否持有，卖出的次数
	dp := make([][2]int, len(prices))

	// 初始状态
	// 第1天的状态初始化
	// 第一天，没有持有，也就是没有买入
	dp[0][0] = 0

	// 第一天，持有，也就是买入股票
	dp[0][1] = -prices[0]

	// 决策，实现状态方程
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])

		if i < 2 {
			dp[i][1] = max(dp[i-1][1], -prices[i])
		}

		if i >= 2 {
			dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i])
		}
	}

	return dp[len(prices)-1][0]
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
