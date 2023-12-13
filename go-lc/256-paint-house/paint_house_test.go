// link		: https://leetcode-cn.com/problems/paint-house/
// Author	: Kylin
// Date		: 2021-12-29

package leetcode

import (
	"fmt"
	"testing"
)

func TestPaintHouse(t *testing.T) {
	ant := [][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}} // 10
	res := MinCost(ant)
	fmt.Println(res)

	bear := [][]int{{7, 6, 2}} // 2
	result := MinCost(bear)
	fmt.Println(result)

}
