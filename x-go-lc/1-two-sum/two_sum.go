// link		: https://leetcode.cn/problems/two-sum
// Author	: Kylin
// Date		: 2023-03-16

package leetcode

// map
func twoSum(nums []int, target int) []int {
	tempMap := make(map[int]int, len(nums))

	for index, num := range nums {
		if nextIndex, ok := tempMap[target-num]; ok {
			return []int{index, nextIndex}
		}
		tempMap[num] = index
	}
	return []int{}
}
