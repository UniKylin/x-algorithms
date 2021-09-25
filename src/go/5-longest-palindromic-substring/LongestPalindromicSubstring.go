// link		: https://leetcode-cn.com/problems/longest-palindromic-substring/
// Author	: Kylin
// Date		: 2021-09-25

func longestPalindrome(s string) string {
	max := 0
	result := ""

	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			tempStr := s[i:j]
			if isPalindromic(tempStr) && len(tempStr) > max {
				result = tempStr
				max = maxium(max, len(tempStr))
			}
		}
	}
	return result
}

func isPalindromic(str string) bool {
	for index := 0; index < len(str)/2; index++ {
		if str[index] != str[len(str)-index-1] {
			return false
		}
	}
	return true
}

func maxium(a, b int) int {
	if a > b {
		return a
	}
	return b
}