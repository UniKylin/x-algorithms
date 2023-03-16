// link		: https://leetcode.cn/problems/longest-substring-without-repeating-characters
// Author	: Kylin
// Date		: 2023-03-16

package leetcode

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	str := "abcabcbb"
	res := lengthOfLongestSubstring(str)
	fmt.Println(res)
}
