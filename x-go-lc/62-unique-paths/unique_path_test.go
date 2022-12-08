// link		: https://leetcode-cn.com/problems/unique-paths/
// Author	: Kylin
// Date		: 2022-01-05

package leetcode

import (
	"fmt"
	"testing"
)

func TestUniquePaths(t *testing.T) {
	ant := UniquePaths(3, 7)
	fmt.Println(">>> m: 3, n: 7 => ", ant)

	bear := UniquePaths(3, 2)
	fmt.Println(">>> m: 3, n: 2 => ", bear)

	cat := UniquePaths(7, 3)
	fmt.Println(">>> m: 7, n: 3 => ", cat)

	dog := UniquePaths(3, 3)
	fmt.Println(">>> m: 3, n: 3 => ", dog)
}
