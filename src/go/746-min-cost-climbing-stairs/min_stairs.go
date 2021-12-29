// link		: https://leetcode-cn.com/problems/min-cost-climbing-stairs/
// Author	: Kylin
// Date		: 2021-12-29

package leetcode

// 思路
// 怎样确定状态转移方程呢
// 台阶可以从 第0层 第一层 开始走，那么从第0层 和 第一层开始的代价都是 0，初始化数据搞定了
// 那么到 N 层台阶其实就有两种可能一个是 N - 2 层 和 N - 1 层 一次走两步 一次走一步上去了
// 然后加上到达N层的代价
// 完事，开干
// 不过这种解法 空间复杂度和时间复杂度都是 O(N)
// 有没有更猛的方式
// 有，当然有
// 可以把状态方程用 cost 的方式来搞定，是不是节省空间开销了
// 其实到这里还是有方法可以继续优化的
// 可以发现转移方程只是利用到了最近的两个数字
// 所以其实动态转移方程的大小是 2
// 就可以把空间复杂度优化到 O(1)

func MinCostClimbingStairs(cost []int) int {
	n := len(cost)

	for step := 2; step < n; step++ {
		cost[step] = min(cost[step-2], cost[step-1]) + cost[step]
	}

	return min(cost[n-2], cost[n-1])
}

func min(ant, bear int) int {
	if ant > bear {
		return bear
	}
	return ant
}
