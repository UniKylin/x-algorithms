// link		: https://leetcode-cn.com/problems/design-linked-list/
// Author	: Kylin
// Date		: 2022-02-18

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

type MyLinkedList struct {
	size int
	head *ListNode
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{0, &ListNode{}}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (l *MyLinkedList) Get(index int) int {
	if index < 0 || index >= l.size {
		return -1
	}

	prev := l.head

	for i := 0; i < index; i++ {
		prev = prev.Next
	}

	return prev.Next.Val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (l *MyLinkedList) AddAtHead(val int) {
	l.AddAtIndex(0, val)
}

/** Append a node of value val to the last element of the linked list. */
func (l *MyLinkedList) AddAtTail(val int) {
	l.AddAtIndex(l.size, val)
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > l.size {
		return
	}

	prev := l.head

	for i := 0; i < index; i++ {
		prev = prev.Next
	}

	node := &ListNode{Val: val}
	node.Next = prev.Next
	prev.Next = node

	l.size++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= l.size {
		return
	}

	prev := l.head

	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	prev.Next = prev.Next.Next

	l.size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
