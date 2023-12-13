// link		: https://leetcode-cn.com/problems/design-circular-deque/
// Author	: Kylin
// Date		: 2022-02-21

package leetcode

import (
	"fmt"
	"testing"
)

func TestCircularDeque(t *testing.T) {
	deque := Constructor(3)

	deque.InsertLast(1)
	deque.InsertLast(2)

	deque.InsertFront(3)
	deque.InsertFront(4)

	fmt.Println(">>> Deque is full:", deque.IsFull())

	rear := deque.GetRear()
	fmt.Println(">>> Rear ele:", rear)

	deque.DeleteLast()
	deque.InsertFront(8)

	frontEle := deque.GetFront()
	fmt.Println(">>> Front ele:", frontEle)
}
