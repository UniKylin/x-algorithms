// link		: https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree/
// Author	: Kylin
// Date		: 2022-03-01

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head
	var preSlow *ListNode = nil

	for fast != nil && fast.Next != nil {
		preSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	root := &TreeNode{Val: slow.Val}
	if preSlow != nil {
		preSlow.Next = nil
		root.Left = sortedListToBST(head)
	}
	root.Right = sortedListToBST(slow.Next)
	return root
}
