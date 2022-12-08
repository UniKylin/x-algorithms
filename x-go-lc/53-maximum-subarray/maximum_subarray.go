// link		: https://leetcode-cn.com/problems/maximum-subarray/
// Author	: Kylin
// Date		: 2022-03-03

package leetcode

func MaxSubArray(nums []int) int {
	max := nums[0]
	lenNum := len(nums)
	for i := 1; i < lenNum; i++ {
		if nums[i] < nums[i]+nums[i-1] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
