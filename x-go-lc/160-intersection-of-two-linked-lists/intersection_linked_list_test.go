// link		: https://leetcode-cn.com/problems/intersection-of-two-linked-lists/
// Author	: Kylin
// Date		: 2022-02-18

package leetcode

import (
	"fmt"
	"testing"
)

func TestInsectionLinkedList(t *testing.T) {
	cat := &ListNode{Val: 4, Next: nil}
	bear := &ListNode{Val: 3, Next: cat}
	ant := &ListNode{Val: 2, Next: bear}
	headA := &ListNode{Val: 1, Next: ant}

	dog := &ListNode{Val: 4, Next: bear}
	headB := &ListNode{Val: 3, Next: dog}

	cur := GetIntersectionNode(headA, headB)

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
