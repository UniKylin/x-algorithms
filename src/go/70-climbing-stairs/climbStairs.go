// link		: https://leetcode-cn.com/problems/climbing-stairs/
// Author	: Kylin
// Date		: 2021-08-28
package leetcode

func climbStairs(n int) int {

	if n <= 2 {
		return n
	}

	var ant int = 1
	var bear int = 2
	var cat int = 3

	for i := 3; i < n+1; i++ {
		cat = ant + bear
		ant = bear
		bear = cat
	}

	return cat
}
