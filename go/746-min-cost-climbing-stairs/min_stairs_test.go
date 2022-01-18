// link		: https://leetcode-cn.com/problems/min-cost-climbing-stairs/
// Author	: Kylin
// Date		: 2021-12-29

package leetcode

import (
	"fmt"
	"testing"
)

func TestMinCostClimbingStairs(t *testing.T) {
	antCost := []int{10, 15, 20}                          // 15
	bearCost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1} //6

	antRes := MinCostClimbingStairs(antCost)
	bearRes := MinCostClimbingStairs(bearCost)

	fmt.Println(">>>", antRes)
	fmt.Println(">>>", bearRes)
}
