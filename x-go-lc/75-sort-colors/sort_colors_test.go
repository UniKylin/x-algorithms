// link		: https://leetcode.cn/problems/sort-colors
// Author	: Kylin
// Date		: 2023-03-12

package leetcode

import (
	"fmt"
	"testing"
)

func TestSortColors(t *testing.T) {
	nums := []int{2, 0, 2, 1, 1, 0}
	SortColors(nums)
	fmt.Println(nums)
}
