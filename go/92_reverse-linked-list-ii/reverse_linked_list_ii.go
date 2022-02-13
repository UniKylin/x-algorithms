// link		: https://leetcode-cn.com/problems/reverse-linked-list-ii/
// Author	: Kylin
// Date		: 2022-02-14

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseLinkedList(head *ListNode) {
	var pre *ListNode
	current := head

	for current != nil {
		next := current.Next
		current.Next = pre
		pre = current
		current = next
	}
}

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{-1, head}
	pre := dummy

	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	rightNode := pre
	for k := 0; k < right-left+1; k++ {
		rightNode = rightNode.Next
	}

	leftNode := pre.Next
	last := rightNode.Next

	pre.Next = nil
	rightNode.Next = nil

	ReverseLinkedList(leftNode)

	pre.Next = rightNode
	leftNode.Next = last

	return dummy.Next
}
