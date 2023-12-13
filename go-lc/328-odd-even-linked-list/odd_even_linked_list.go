// link		: https://leetcode-cn.com/problems/odd-even-linked-list/
// Author	: Kylin
// Date		: 2022-02-18

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func OddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	evenHead := head.Next
	odd := head
	even := evenHead

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = evenHead

	return head

}
