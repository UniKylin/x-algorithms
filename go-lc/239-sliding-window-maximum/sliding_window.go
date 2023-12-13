// link		: https://leetcode.cn/problems/sliding-window-maximum
// Author	: Kylin
// Date		: 2022-12-31

package leetcode

func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	var queue []int

	// 遍历数组开干
	for index, val := range nums {
		// 和队列尾部数进行对比
		for len(queue) > 0 && val >= nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}

		// 存储的是数组索引
		queue = append(queue, index)

		if index-k+1 > queue[0] {
			queue = queue[1:]
		}

		if index+1 >= k {
			res = append(res, nums[queue[0]])
		}
	}

	return res
}
