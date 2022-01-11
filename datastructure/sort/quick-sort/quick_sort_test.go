package sort

import (
	"fmt"
	"testing"
)

func TestQuickSrot(t *testing.T) {
	arr := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	res := QuickSort(arr)
	fmt.Println(res)
}
