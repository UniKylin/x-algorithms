// link		: https://leetcode-cn.com/problems/sqrtx/
// Author	: Kylin
// Date		: 2022-01-16

package leetcode

func Sqrtx(x int) int {
	left, right := 0, x
	res := -1

	for left <= right {
		mid := left + (right-left)>>1

		if mid*mid <= x {
			res = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return res
}
