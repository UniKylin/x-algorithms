// link		: https://leetcode-cn.com/problems/insertion-sort-list/
// Author	: Kylin
// Date		: 2022-02-10

package leetcode

import (
	"fmt"
	"testing"
)

func TestInsertSortLst(t *testing.T) {
	ant3 := &ListNode{Val: 5, Next: nil}
	ant2 := &ListNode{Val: 1, Next: ant3}
	ant1 := &ListNode{Val: 2, Next: ant2}
	ant := &ListNode{Val: 4, Next: ant1}

	res := InsertionSortList(ant)

	for res != nil {
		if res.Next != nil {
			fmt.Print(res.Val, " -> ")
		} else {
			fmt.Println(res.Val)
		}
		res = res.Next
	}
}
