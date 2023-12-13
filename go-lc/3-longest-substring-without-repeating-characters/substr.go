// link		: https://leetcode.cn/problems/longest-substring-without-repeating-characters
// Author	: Kylin
// Date		: 2023-03-16

package leetcode

func lengthOfLongestSubstring(s string) (ans int) {
	cnt := [128]int{}
	l, r := 0, 0
	for r < len(s) {
		for cnt[s[r]] > 0 {
			cnt[s[l]]--
			l++
		}
		cnt[s[r]]++
		ans = max(ans, r-l+1)
		r++
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
