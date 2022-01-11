package sort

import (
	"fmt"
	"testing"
)

func TestQuickSrotInPlace(t *testing.T) {
	arr := []int{5, 1, 3, 9, 7, 2, 8, 4, 6}
	res := QuickSort(arr)
	fmt.Println(res)
}
