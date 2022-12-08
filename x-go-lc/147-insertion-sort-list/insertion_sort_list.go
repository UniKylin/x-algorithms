// link		: https://leetcode-cn.com/problems/insertion-sort-list/
// Author	: Kylin
// Date		: 2022-02-10

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func InsertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head}
	lastSorted, cur := head, head.Next

	for cur != nil {
		if lastSorted.Val <= cur.Val {
			lastSorted = lastSorted.Next
		} else {
			prev := dummy
			for prev.Next.Val <= cur.Val {
				prev = prev.Next
			}

			lastSorted.Next = cur.Next
			cur.Next = prev.Next
			prev.Next = cur
		}
		cur = lastSorted.Next
	}

	return dummy.Next
}
