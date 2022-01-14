// link		: https://leetcode-cn.com/problems/binary-search/
// Author	: Kylin
// Date		: 2022-01-14

package leetcode

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	res := Search(nums, 9)
	fmt.Println(res)

	nextRes := Search(nums, 2)
	fmt.Println(nextRes)
}
