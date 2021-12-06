package leetcode

import (
	"fmt"
	"testing"
)

func TestSellStock(t *testing.T) {
	ant := []int{1, 2, 3, 0, 2}
	profit := MaxProfitCooldown(ant)
	fmt.Println(profit)
}
