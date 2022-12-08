// link		: https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
// Author	: Kylin
// Date		: 2022-02-15

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteDuplicates(head *ListNode) *ListNode {
	cur := head

	for cur != nil {
		if cur.Next != nil && cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return head
}
