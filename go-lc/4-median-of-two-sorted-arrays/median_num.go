// link		: https://leetcode.cn/problems/median-of-two-sorted-arrays
// Author	: Kylin
// Date		: 2023-03-16

package leetcode

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mid := (len(nums1) + len(nums2) + 1) / 2
	count := 0
	for i, j := 0, 0; i < len(nums1) || j < len(nums2); {
		a, b := math.MaxInt32, math.MaxInt32
		if i < len(nums1) {
			a = nums1[i]
		}
		if j < len(nums2) {
			b = nums2[j]
		}
		var t int
		if a < b {
			t = nums1[i]
			i++
		} else {
			t = nums2[j]
			j++
		}
		count++
		if count == mid {
			if (len(nums1)+len(nums2))%2 == 1 {
				return float64(t)
			} else {
				a, b = math.MaxInt32, math.MaxInt32
				if i < len(nums1) {
					a = nums1[i]
				}
				if j < len(nums2) {
					b = nums2[j]
				}
				if a > b {
					a = b
				}
				return float64(a+t) / 2.0
			}
		}
	}
	return 0
}
