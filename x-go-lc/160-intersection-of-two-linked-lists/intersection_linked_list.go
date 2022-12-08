// link		: https://leetcode-cn.com/problems/intersection-of-two-linked-lists/
// Author	: Kylin
// Date		: 2022-02-18

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	pA, pB := headA, headB

	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}

		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}

	return pA
}
