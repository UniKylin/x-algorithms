// link		: https://leetcode-cn.com/problems/counting-bits/
// Author	: Kylin
// Date		: 2022-02-02

package leetcode

func CountBits(n int) []int {
	bits := make([]int, n+1)

	// 1111 -> 15
	// 1110 -> 14
	// 1100 -> 12
	// 1000 -> 8
	// 0000 -> 0
	// x := 15
	// y := x & (x - 1)
	// fmt.Println(y)

	for i := 1; i <= n; i++ {
		bits[i] = bits[i&(i-1)] + 1
	}

	return bits
}
