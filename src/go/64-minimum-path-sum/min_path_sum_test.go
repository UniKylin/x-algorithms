// link		: https://leetcode-cn.com/problems/minimum-path-sum/
// Author	: Kylin
// Date		: 2022-01-05

package leetcode

import (
	"fmt"
	"testing"
)

func TestMinPathSum(t *testing.T) {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	res := MinPathSum(grid)
	fmt.Println(">>> res:", res) // 7

	gridNext := [][]int{{1, 2, 3}, {4, 5, 6}}
	resNext := MinPathSum(gridNext)
	fmt.Println(">>> next res:", resNext) // 12
}
