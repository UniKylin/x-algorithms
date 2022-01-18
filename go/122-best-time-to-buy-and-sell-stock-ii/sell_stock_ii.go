// link		: https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
// Author	: Kylin
// Date		: 2021-12-06

package leetcode

func MaxProfitII(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	dp := make([][2]int, len(prices))

	dp[0][0] = 0
	dp[0][1] = -prices[0] // 可以理解为第一天买入是亏损的

	for i := 1; i < len(prices); i++ {
		// 不持有 那么在 i - 1 天无交易，在 i 卖出获取收益
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])

		// 持有 那么在 i - 1 天无交易，在 i 天继续买入
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[len(prices)-1][0]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
