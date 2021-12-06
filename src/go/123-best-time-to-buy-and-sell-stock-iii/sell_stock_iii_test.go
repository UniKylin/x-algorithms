package leetcode

import (
	"fmt"
	"testing"
)

func TestSellStock(t *testing.T) {
	ant := []int{3, 3, 5, 0, 0, 3, 1, 4}
	profit := MaxProfitIII(ant)
	fmt.Println(profit)

	bear := []int{1, 2, 3, 4, 5}
	nextProfit := MaxProfitIII(bear)
	fmt.Println(nextProfit)

	cat := []int{7, 6, 4, 3, 1}
	lastProfit := MaxProfitIII(cat)
	fmt.Println(lastProfit)

	dog := []int{1}
	finalProfit := MaxProfitIII(dog)
	fmt.Println(finalProfit)
}
