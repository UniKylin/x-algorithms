// link		: https://leetcode.cn/problems/median-of-two-sorted-arrays
// Author	: Kylin
// Date		: 2023-03-16

package leetcode

import (
	"fmt"
	"testing"
)

func TestMedianNum(t *testing.T) {
	ant := []int{1, 2}
	bear := []int{3, 4}
	res := findMedianSortedArrays(ant, bear)
	fmt.Println(res)
}
