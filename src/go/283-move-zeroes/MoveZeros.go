// link		: https://leetcode-cn.com/problems/move-zeroes/
// Author	: Kylin
// Date		: 2021-09-02
package leetcode

func moveZeroes(nums []int) {
	slow := 0
	fast := 0

	for ; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}

		if fast == len(nums)-1 {
			for slow != len(nums) {
				nums[slow] = 0
				slow++
			}
		}
	}
}
