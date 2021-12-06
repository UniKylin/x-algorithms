// link		: https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/
// Author	: Kylin
// Date		: 2021-12-06

package leetcode

import "math"

func MaxProfitIII(prices []int) int {
	buyFirst, buyNext := math.MaxInt32, math.MaxInt32
	sellFirst, sellNext := math.MinInt32, math.MinInt32

	for i := 0; i < len(prices); i++ {
		buyFirst = min(buyFirst, prices[i])
		sellFirst = max(sellFirst, prices[i]-buyFirst)

		// 第二次购买的价格 = 股票价格 - 第一次的收益
		buyNext = min(buyNext, prices[i]-sellFirst)
		sellNext = max(sellNext, prices[i]-buyNext)
	}

	return sellNext
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
