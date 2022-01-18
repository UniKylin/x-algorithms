package leetcode

import (
	"fmt"
	"testing"
)

func TestSellStock(t *testing.T) {
	ant := []int{1, 3, 2, 8, 4, 9}
	profit := MaxProfitWithFee(ant, 2)
	fmt.Println(profit)

	bear := []int{1, 3, 7, 5, 10, 3}
	nextProfit := MaxProfitWithFee(bear, 3)
	fmt.Println(nextProfit)
}
