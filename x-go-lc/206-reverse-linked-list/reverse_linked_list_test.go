// link		: https://leetcode-cn.com/problems/reverse-linked-list/
// Author	: Kylin
// Date		: 2022-01-20

package leetcode

import (
	"fmt"
	"testing"
)

func TestReverseList(test *testing.T) {
	node5 := &ListNode{6, nil}
	node4 := &ListNode{5, node5}
	node3 := &ListNode{4, node4}
	node2 := &ListNode{3, node3}
	node1 := &ListNode{2, node2}
	node := &ListNode{1, node1}

	cur := ReverseList(node)

	for ; cur != nil; cur = cur.Next {
		if cur.Next != nil {
			fmt.Print(cur.Val, " -> ")
		} else {
			fmt.Print(cur.Val)
		}
	}
	fmt.Println()
}
