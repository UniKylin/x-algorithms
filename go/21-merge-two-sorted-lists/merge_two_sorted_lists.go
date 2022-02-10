// link		: https://leetcode-cn.com/problems/merge-two-sorted-lists/
// Author	: Kylin
// Date		: 2022-02-10

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(ant *ListNode, bear *ListNode) *ListNode {
	dummy := &ListNode{}
	prev := dummy

	for ant != nil && bear != nil {
		if ant.Val < bear.Val {
			prev.Next = ant
			ant = ant.Next
		} else {
			prev.Next = bear
			bear = bear.Next
		}
		prev = prev.Next
	}

	if ant == nil {
		prev.Next = bear
	} else {
		prev.Next = ant
	}

	return dummy.Next
}
