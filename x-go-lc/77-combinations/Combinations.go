// link		: https://leetcode-cn.com/problems/combinations/
// Author	: Kylin
// Date		: 2021-10-13

package leetcode

func combine(n int, k int) [][]int {
	var subset []int
	var result [][]int

	var dfs func(index int)
	dfs = func(index int) {
		if len(subset) == k {
			result = append(result, append([]int{}, subset...))
			return
		}

		if index > n {
			return
		}

		dfs(index + 1)
		subset = append(subset, index)
		dfs(index + 1)
		subset = subset[:len(subset)-1]
	}

	// Notice 1....n
	dfs(1)
	return result
}
