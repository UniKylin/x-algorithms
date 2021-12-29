// link		: https://leetcode-cn.com/problems/house-robber/
// Author	: Kylin
// Date		: 2021-12-29

package leetcode

func Robber(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	slow := nums[0]
	next := max(nums[0], nums[1])

	for index := 2; index < n; index++ {
		slow, next = next, max(slow+nums[index], next)
	}
	return next
}

func max(ant, bear int) int {
	if ant > bear {
		return ant
	}
	return bear
}
