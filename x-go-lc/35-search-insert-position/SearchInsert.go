// link		: https://leetcode-cn.com/problems/search-insert-position/
// Author	: Kylin
// Date		: 2021-09-02

package leetcode

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums)

	for left < right {
		middle := left + (right-left)>>1

		if target > nums[middle] {
			left = middle + 1
		} else {
			right = middle
		}
	}

	return left
}
