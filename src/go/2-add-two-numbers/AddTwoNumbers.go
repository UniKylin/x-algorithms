// link		: https://leetcode-cn.com/problems/add-two-numbers/
// Author	: Kylin
// Date		: 2021-09-05

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(ant *ListNode, bear *ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode
	carry := 0

	for ant != nil || bear != nil {
		n1, n2 := 0, 0

		if ant != nil {
			n1 = ant.Val
			ant = ant.Next
		}

		if bear != nil {
			n2 = bear.Val
			bear = bear.Next
		}

		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10

		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}

	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}

	return head

}
