// link		: https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
// Author	: Kylin
// Date		: 2022-02-15

package leetcode

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicate(t *testing.T) {
	dog := &ListNode{Val: 3, Next: nil}
	cat := &ListNode{Val: 2, Next: dog}
	bear := &ListNode{Val: 2, Next: cat}
	ant := &ListNode{Val: 1, Next: bear}
	head := &ListNode{Val: 1, Next: ant}

	cur := DeleteDuplicates(head)

	for cur != nil {
		if cur.Next != nil {
			fmt.Print(cur.Val, "->")
		} else {
			fmt.Print(cur.Val)
		}
		cur = cur.Next
	}
	fmt.Println()
}
