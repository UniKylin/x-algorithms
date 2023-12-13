// link		: https://leetcode.cn/problems/sort-colors
// Author	: Kylin
// Date		: 2023-03-12

package leetcode

func SortColors(nums []int) {
	zero := -1
	two := len(nums)

	for i := 0; i < two; {
		if nums[i] == 1 {
			i++
		} else if nums[i] == 2 {
			two--
			nums[i], nums[two] = nums[two], nums[i]
		} else {
			zero++
			nums[i], nums[zero] = nums[zero], nums[i]
			i++
		}
	}
}
