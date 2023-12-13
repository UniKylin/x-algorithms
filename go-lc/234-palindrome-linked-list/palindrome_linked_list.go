// link		: https://leetcode-cn.com/problems/palindrome-linked-list/
// Author	: Kylin
// Date		: 2022-02-17

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func IsPalindrome(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return true
	}

	var pre *ListNode
	var temp *ListNode
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		temp = slow.Next
		slow.Next = pre
		pre = slow
		slow = temp
	}

	if fast != nil {
		slow = slow.Next
	}

	for pre != nil && slow != nil {
		if pre.Val != slow.Val {
			return false
		}
		pre = pre.Next
		slow = slow.Next
	}

	return true
}
