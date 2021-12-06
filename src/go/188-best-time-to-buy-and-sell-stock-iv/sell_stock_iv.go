// link		: https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/
// Author	: Kylin
// Date		: 2021-12-06

package leetcode

import "math"

func MaxProfitIV(k int, prices []int) int {
	n := len(prices)
	if k == 0 || n == 0 {
		return 0
	}
	buy := make([]int, k)
	sell := make([]int, k)
	for i := 0; i < k; i++ {
		buy[i] = -prices[0]
	}

	for i := 1; i < n; i++ {
		buy[0] = int(math.Max(float64(buy[0]), float64(-prices[i])))
		sell[0] = int(math.Max(float64(sell[0]), float64(buy[0]+prices[i])))

		for j := 1; j < k; j++ {
			buy[j] = int(math.Max(float64(buy[j]), float64(sell[j-1]-prices[i])))
			sell[j] = int(math.Max(float64(sell[j]), float64(buy[j]+prices[i])))
		}
	}

	return sell[k-1]
}
