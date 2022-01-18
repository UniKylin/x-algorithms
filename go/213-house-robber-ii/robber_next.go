// link		: https://leetcode-cn.com/problems/house-robber-ii/
// Author	: Kylin
// Date		: 2021-12-29

package leetcode

// 如果只有一间直接放心偷
// 如果有两间那么只能偷一间
// 如果有两间以上，因为 第N间和第0间房屋是挨着的，所以存在两种情况
// 第一种：抢第一家不抢最后一家
// 第二种: 抢最后一家不抢第一家
func RobberNext(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	if n == 2 {
		return max(nums[0], nums[1])
	}

	first := work(nums, 0, n-1)
	last := work(nums, 1, n)

	return max(first, last)
}

func work(nums []int, start int, end int) int {
	// 优化后的空间复杂度
	// 两个指针向后挪动，如果是 Java就需要一个中间变量
	first := nums[start]
	last := max(nums[start], nums[start+1])

	for i := start + 2; i < end; i++ {
		first, last = last, max(first+nums[i], last)
	}

	return last
}

func max(ant, bear int) int {
	if ant > bear {
		return ant
	}
	return bear
}
