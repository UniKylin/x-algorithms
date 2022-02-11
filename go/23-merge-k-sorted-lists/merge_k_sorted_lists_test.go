// link		: https://leetcode-cn.com/problems/merge-two-sorted-lists/
// Author	: Kylin
// Date		: 2022-02-11

package leetcode

import (
	"fmt"
	"testing"
)

func TestMergeKSortedLits(t *testing.T) {
	ant2 := &ListNode{Val: 8, Next: nil}
	ant1 := &ListNode{Val: 3, Next: ant2}
	ant := &ListNode{Val: 1, Next: ant1}

	bear2 := &ListNode{Val: 4, Next: nil}
	bear1 := &ListNode{Val: 2, Next: bear2}
	bear := &ListNode{Val: 1, Next: bear1}

	cat1 := &ListNode{Val: 6, Next: nil}
	cat := &ListNode{Val: 5, Next: cat1}

	lists := []*ListNode{ant, bear, cat}

	res := MergeKLists(lists)

	for res != nil {
		if res.Next != nil {
			fmt.Print(res.Val, " -> ")
		} else {
			fmt.Println(res.Val)
		}
		res = res.Next
	}
}
