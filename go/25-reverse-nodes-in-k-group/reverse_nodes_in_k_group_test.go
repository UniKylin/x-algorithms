// link		: https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
// Author	: Kylin
// Date		: 2022-02-20

package leetcode

import (
	"fmt"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	dog := &ListNode{Val: 12, Next: nil}
	cat := &ListNode{Val: 8, Next: dog}
	bear := &ListNode{Val: 3, Next: cat}
	ant := &ListNode{Val: 1, Next: bear}

	res := ReverseKGroup(ant, 2)

	for res != nil {
		if res.Next != nil {
			fmt.Print(res.Val, " -> ")
		} else {
			fmt.Println(res.Val)
		}
		res = res.Next
	}
}
