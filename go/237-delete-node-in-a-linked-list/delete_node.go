// link		: https://leetcode-cn.com/problems/delete-node-in-a-linked-list/
// Author	: Kylin
// Date		: 2022-02-13

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
