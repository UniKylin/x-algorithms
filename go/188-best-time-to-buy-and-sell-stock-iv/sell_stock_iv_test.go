package leetcode

import (
	"fmt"
	"testing"
)

func TestSellStock(t *testing.T) {
	ant := []int{2, 4, 1}
	profit := MaxProfitIV(2, ant)
	fmt.Println(profit)

	bear := []int{3, 2, 6, 5, 0, 3}
	nextProfit := MaxProfitIV(2, bear)
	fmt.Println(nextProfit)

}
