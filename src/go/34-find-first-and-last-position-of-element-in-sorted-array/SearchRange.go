// link		: https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
// Author	: Kylin
// Date		: 2021-09-02

func searchRange(nums []int, target int) []int {
	leftIndex := binarySearchLeft(nums, target)
	rightIndex := binarySearchRight(nums, target)

	if leftIndex <= rightIndex && rightIndex < len(nums)
		&& nums[leftIndex] == target && nums[rightIndex] == target {
		return []int{leftIndex, rightIndex}
	}

	return []int{-1, -1}
}

func binarySearchLeft(nums []int, target int) int {
	left := -1
	right := len(nums)

	for left+1 != right {
		middle := (left + right) >> 1

		if nums[middle] >= target {
			right = middle
		} else {
			left = middle
		}
	}

	return right
}

func binarySearchRight(nums []int, target int) int {
	left := -1
	right := len(nums)

	for left+1 != right {
		middle := (left + right) >> 1

		if nums[middle] <= target {
			left = middle
		} else {
			right = middle
		}
	}

	return left
}