// link		: https://leetcode-cn.com/problems/palindrome-linked-list/
// Author	: Kylin
// Date		: 2022-02-17

package leetcode

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicate(t *testing.T) {
	cat := &ListNode{Val: 1, Next: nil}
	bear := &ListNode{Val: 2, Next: cat}
	ant := &ListNode{Val: 2, Next: bear}
	head := &ListNode{Val: 1, Next: ant}

	fmt.Println(IsPalindrome(head))
}
