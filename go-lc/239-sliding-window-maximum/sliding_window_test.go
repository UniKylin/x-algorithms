// link		: https://leetcode.cn/problems/sliding-window-maximum
// Author	: Kylin
// Date		: 2022-12-31

package leetcode

import (
	"fmt"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 2
	res := maxSlidingWindow(nums, k)

	for _, data := range res {
		fmt.Print(data, "\t")
	}
}
