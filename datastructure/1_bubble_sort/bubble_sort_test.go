package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{2, 1, 4, 6, 3, 9, 7}
	BubbleSort(arr)
	fmt.Println(arr)

	list := []int{1, 2, 3, 4, 5, 6, 7}
	BubbleSort(list)
	fmt.Println(list)
}
