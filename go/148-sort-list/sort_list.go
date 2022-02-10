// link		: https://leetcode-cn.com/problems/sort-list/
// Author	: Kylin
// Date		: 2022-02-10

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func SortList(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	cur := head

	var prev, temp *ListNode

	for cur != nil && cur.Next != nil {
		if cur.Val <= cur.Next.Val {
			cur = cur.Next
		} else {
			temp = cur.Next

			cur.Next = cur.Next.Next

			prev = dummy
			for prev.Next.Val <= temp.Val {
				prev = prev.Next
			}

			temp.Next = prev.Next
			prev.Next = temp
		}
	}

	return dummy.Next
}
