package sort

func BubbleSort(arr []int) {
	n := len(arr)

	swap := false

	for i := n - 1; i > 0; i-- {
		// 0 1 ... n-3 n-2
		// 0 1 ... n-3
		// 0 1
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swap = true
			}
		}

		if !swap {
			break
		}
	}
}
