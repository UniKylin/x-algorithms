package sort

import "fmt"

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var left, right []int

	for _, ele := range arr[1:] {
		if ele <= pivot {
			left = append(left, ele)
		} else {
			right = append(right, ele)
		}
	}
	fmt.Println(">>> pivot:", pivot)
	fmt.Println(">>> arr:", arr)

	return append(QuickSort(left), append([]int{pivot}, QuickSort(right)...)...)
}
