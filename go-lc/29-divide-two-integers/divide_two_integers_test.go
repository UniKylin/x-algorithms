// link		: https://leetcode-cn.com/problems/divide-two-integers/
// Author	: Kylin
// Date		: 2022-02-02

package leetcode

import (
	"fmt"
	"testing"
)

func TestDivide(t *testing.T) {
	dividend := 15
	divisor := 3

	res := Divide(dividend, divisor)
	fmt.Println(">>> res:", res)
}
