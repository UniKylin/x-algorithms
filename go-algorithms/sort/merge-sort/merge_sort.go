package sort

import "fmt"

func MergeSort(arr []int) []int {
	n := len(arr)

	if n == 1 {
		return arr
	}

	m := n / 2

	leftArr := MergeSort(arr[:m])
	rightArr := MergeSort(arr[m:])

	return merge(leftArr, rightArr)
}

func merge(l []int, r []int) []int {
	fmt.Println(">>>l:", l)
	fmt.Println(">>>r:", r)

	leftLen := len(l)
	rightLen := len(r)

	res := make([]int, 0)

	leftIndex, rightIndex := 0, 0

	for leftIndex < leftLen && rightIndex < rightLen {
		if l[leftIndex] > r[rightIndex] {
			res = append(res, r[rightIndex])
			rightIndex++
		} else {
			res = append(res, l[leftIndex])
			leftIndex++
		}
	}

	if leftIndex < leftLen {
		res = append(res, l[leftIndex:]...)
	}

	if rightIndex < rightLen {
		res = append(res, r[rightIndex:]...)
	}

	return res
}
