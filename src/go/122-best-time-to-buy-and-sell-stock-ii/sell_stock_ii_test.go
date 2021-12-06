package leetcode

import (
	"fmt"
	"testing"
)

func TestSellStock(t *testing.T) {
	ant := []int{7, 1, 5, 3, 6, 4}
	profit := MaxProfitII(ant)
	fmt.Println(profit)

	bear := []int{1, 2, 3, 4, 5}
	nextProfit := MaxProfitII(bear)
	fmt.Println(nextProfit)

	cat := []int{7, 6, 4, 3, 1}
	lastProfit := MaxProfitII(cat)
	fmt.Println(lastProfit)
}
