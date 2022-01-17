// link		: https://leetcode-cn.com/problems/design-circular-queue/
// Author	: Kylin
// Date		: 2022-01-17

package leetcode

type Node struct {
	value int
	next  *Node
}

type MyCircularQueue struct {
	capacity int
	length   int
	head     *Node
	tail     *Node
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		capacity: k,
		length:   0,
	}
}

func (c *MyCircularQueue) EnQueue(val int) bool {
	if c.length+1 > c.capacity {
		return false
	}

	node := Node{value: val}

	if c.length == 0 {
		c.head = &node
		c.tail = &node
	} else {
		c.tail.next = &node
		c.tail = c.tail.next
	}
	c.length++
	return true
}

func (c *MyCircularQueue) DeQueue() bool {
	if c.length == 0 {
		return false
	}

	c.head = c.head.next
	c.length--
	return true
}

func (c *MyCircularQueue) Front() int {
	if c.length == 0 {
		return -1
	}

	return c.head.value
}

func (c *MyCircularQueue) Rear() int {
	if c.length == 0 {
		return -1
	}

	return c.tail.value
}

func (c *MyCircularQueue) IsEmpty() bool {
	return c.length == 0
}

func (c *MyCircularQueue) IsFull() bool {
	return c.length == c.capacity
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
