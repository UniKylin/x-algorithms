// link		: https://leetcode-cn.com/problems/split-linked-list-in-parts/
// Author	: Kylin
// Date		: 2021-09-25

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(head *ListNode, k int) []*ListNode {
	count := 0

	for node := head; node != nil; node = node.Next {
		count++
	}

	part, remain := count/k, count%k

	parts := make([]*ListNode, k)

	for i, current := 0, head; i < k && current != nil; i++ {
		parts[i] = current
		partSize := part

		if i < remain {
			partSize++
		}

		for j := 1; j < partSize; j++ {
			current = current.Next
		}
		current, current.Next = current.Next, nil
	}

	return parts
}