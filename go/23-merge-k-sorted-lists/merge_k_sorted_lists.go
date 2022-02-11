// link		: https://leetcode-cn.com/problems/merge-two-sorted-lists/
// Author	: Kylin
// Date		: 2022-02-11

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	res := lists[0]

	for i := 1; i < len(lists); i++ {
		res = merge(res, lists[i])
	}

	return res
}

func merge(ant, bear *ListNode) *ListNode {
	if ant == nil {
		return bear
	}

	if bear == nil {
		return ant
	}

	if ant.Val <= bear.Val {
		ant.Next = merge(ant.Next, bear)
		return ant
	} else {
		bear.Next = merge(ant, bear.Next)
		return bear
	}
}
