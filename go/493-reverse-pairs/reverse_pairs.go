// link		: https://leetcode-cn.com/problems/reverse-pairs/
// Author	: Kylin
// Date		: 2022-02-12

package leetcode

func ReversePairs(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	n1 := append([]int(nil), nums[:n/2]...)
	n2 := append([]int(nil), nums[n/2:]...)
	cnt := reversePairs(n1) + reversePairs(n2)

	j := 0
	for _, v := range n1 {
		for j < len(n2) && v > 2*n2[j] {
			j++
		}
		cnt += j
	}

	p1, p2 := 0, 0
	for i := range nums {
		if p1 < len(n1) && (p2 == len(n2) || n1[p1] <= n2[p2]) {
			nums[i] = n1[p1]
			p1++
		} else {
			nums[i] = n2[p2]
			p2++
		}
	}
	return cnt
}
