// link		: https://leetcode-cn.com/problems/subsets/
// Author	: Kylin
// Date		: 2021-10-13

package leetcode

func subsets(nums []int) [][]int {
	total := len(nums)
	var subset []int
	var result [][]int

	var dfs func(index int)
	dfs = func(index int) {
		if index == total {
			result = append(result, append([]int{}, subset...))
			return
		}

		dfs(index + 1)
		subset = append(subset, nums[index])
		dfs(index + 1)
		subset = subset[:len(subset)-1]
	}

	dfs(0)
	return result
}
