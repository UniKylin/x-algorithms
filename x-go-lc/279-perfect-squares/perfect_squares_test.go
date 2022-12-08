// link		: https://leetcode-cn.com/problems/perfect-squares/
// Author	: Kylin
// Date		: 2022-01-19

package leetcode

import (
	"fmt"
	"testing"
)

func TestNumSquare(t *testing.T) {
	n := 12
	res := NumSquares(n)
	fmt.Println(">>> res:", res)

	n1 := 13
	res1 := NumSquares(n1)
	fmt.Println(">>> res1:", res1)
}
