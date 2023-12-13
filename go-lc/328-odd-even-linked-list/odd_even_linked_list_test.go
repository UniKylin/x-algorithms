// link		: https://leetcode-cn.com/problems/odd-even-linked-list/
// Author	: Kylin
// Date		: 2022-02-18

package leetcode

import (
	"fmt"
	"testing"
)

func TestOddEvenLinkedList(t *testing.T) {
	cat := &ListNode{Val: 4, Next: nil}
	bear := &ListNode{Val: 3, Next: cat}
	ant := &ListNode{Val: 2, Next: bear}
	head := &ListNode{Val: 1, Next: ant}

	cur := OddEvenList(head)

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
