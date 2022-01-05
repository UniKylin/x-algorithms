// link		: https://leetcode-cn.com/problems/unique-paths/
// Author	: Kylin
// Date		: 2022-01-05

package leetcode

func UniquePaths(m int, n int) int {
	dp := make([][]int, m)

	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1 // column
	}

	for j := 0; j < n; j++ {
		dp[0][j] = 1 // init first row
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
