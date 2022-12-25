package basic

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	var arr [5]int
	fmt.Println(">>> arr: ", arr)

	arr[2] = 2009
	fmt.Println(">>> arr: ", arr)
	fmt.Println(">>> arr len: ", len(arr))

	ant := []int{1, 2, 3, 4}
	fmt.Println(">>> arr ant: ", ant)

	// using range traversal
	for index, val := range ant {
		fmt.Println(index, val)
	}

	var bear [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			bear[i][j] = i + j
		}
	}
	fmt.Println(">>> bear :", bear)
}
