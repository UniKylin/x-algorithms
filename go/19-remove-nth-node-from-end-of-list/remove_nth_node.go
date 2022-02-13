// link		: https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
// Author	: Kylin
// Date		: 2022-02-14

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	slow, fast := dummy, head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for ; fast != nil; fast = fast.Next {
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}
