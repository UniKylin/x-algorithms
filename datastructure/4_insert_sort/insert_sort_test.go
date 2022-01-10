package sort

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	arr := []int{9, 1, 4, 8, 7, 3, 7, 2, 5, 6}
	InsertSort(arr)
	fmt.Println(">>>", arr)
}
