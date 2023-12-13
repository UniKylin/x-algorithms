// link		: https://leetcode-cn.com/problems/maximum-subarray/
// Author	: Kylin
// Date		: 2022-03-03

package leetcode

import (
	"fmt"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	res := MaxSubArray(nums)
	fmt.Println(">>> res:", res)
}
