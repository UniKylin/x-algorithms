// link		: https://leetcode-cn.com/problems/rotate-list/
// Author	: Kylin
// Date		: 2022-02-13

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func RotateRight(head *ListNode, k int) *ListNode {

	if k == 0 || head == nil || head.Next == nil {
		return head
	}

	n := 1
	current := head

	for current.Next != nil {
		current = current.Next
		n++
	}

	pointer := n - k%n
	if pointer == n {
		return head
	}

	current.Next = head

	for pointer > 0 {
		current = current.Next
		pointer--
	}

	result := current.Next
	current.Next = nil

	return result

}
