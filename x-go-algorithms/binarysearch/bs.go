package binarysearch

func BinarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left < right {
		middle := left + (right-left)>>1

		if arr[middle] > target {
			right = middle - 1
		} else if arr[middle] < target {
			left = middle + 1
		} else {
			return middle
		}
	}
	return -1
}
