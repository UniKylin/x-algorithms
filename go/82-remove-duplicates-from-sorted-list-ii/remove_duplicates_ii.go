// link		: https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/
// Author	: Kylin
// Date		: 2022-02-15

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{0, head}
	curr := dummy
	for curr.Next != nil && curr.Next.Next != nil {
		if curr.Next.Val == curr.Next.Next.Val {
			x := curr.Next.Val
			for curr.Next != nil && curr.Next.Val == x {
				curr.Next = curr.Next.Next
			}
		} else {
			curr = curr.Next
		}
	}
	return dummy.Next
}
