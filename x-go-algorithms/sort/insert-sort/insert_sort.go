package sort

func InsertSort(arr []int) {
	n := len(arr)

	for i := 1; i < n; i++ {
		pre := i - 1
		cur := arr[i]

		for pre >= 0 && arr[pre] > cur {
			arr[pre+1] = arr[pre]
			pre--
		}

		arr[pre+1] = cur
	}
}
