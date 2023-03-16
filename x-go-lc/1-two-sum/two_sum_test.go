// link		: https://leetcode.cn/problems/two-sum
// Author	: Kylin
// Date		: 2023-03-16

package leetcode

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	ants := []int{2, 7, 11, 15}
	ant := 9
	res := twoSum(ants, ant)
	fmt.Println(res)
}
