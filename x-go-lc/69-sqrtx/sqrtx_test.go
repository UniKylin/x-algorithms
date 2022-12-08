// link		: https://leetcode-cn.com/problems/sqrtx/
// Author	: Kylin
// Date		: 2022-01-16

package leetcode

import (
	"fmt"
	"testing"
)

func TestSqrtx(t *testing.T) {
	res := Sqrtx(4)
	fmt.Println(res)

	next := Sqrtx(27)
	fmt.Println(next)
}
