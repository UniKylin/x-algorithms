package sort

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	arr := []int{9, 1, 4, 7, 3, 7, 2}
	SelectionSort(arr)
	fmt.Println(">>>", arr)

	
}
