package sort

import "fmt"

func QuickSort(arr []int) []int {
	QuickSortSection(arr, 0, len(arr)-1)
	return arr
}

func QuickSortSection(arr []int, low int, high int) {
	if len(arr) <= 1 {
		return
	}

	if low < high {
		pivot := partition(arr, low, high)
		fmt.Println(">>> pivot:", pivot)
		QuickSortSection(arr, low, pivot-1)
		QuickSortSection(arr, pivot+1, high)
	}

}

// Find the pivot and partition the array
func partition(arr []int, low int, high int) int {
	index := low - 1
	pivot := arr[high]

	for i := low; i < high; i++ {
		if arr[i] <= pivot {
			index++
			arr[index], arr[i] = arr[i], arr[index]
			fmt.Println(">>> i: ", i, " arr:", arr)
		}
		fmt.Println()
	}

	arr[index+1], arr[high] = arr[high], arr[index+1]
	fmt.Println("==> arr:", arr)

	return index + 1
}
