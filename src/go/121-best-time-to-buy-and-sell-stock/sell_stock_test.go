package leetcode

import (
	"fmt"
	"testing"
)

func TestSellStock(t *testing.T) {
	ant := []int{7, 1, 5, 3, 6, 4}
	profit := MaxProfit(ant)
	fmt.Println(profit)

	bear := []int{7, 6, 4, 3, 1}
	nextProfit := MaxProfit(bear)
	fmt.Println(nextProfit)
}
