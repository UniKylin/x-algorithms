package binarysearch

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{2, 8, 11, 23, 45, 51, 88, 92}
	target := 23

	res := BinarySearch(arr, target)
	fmt.Println(">>> res:", res)
}
