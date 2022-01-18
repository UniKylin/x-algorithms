// link		: https://leetcode-cn.com/problems/design-circular-queue/
// Author	: Kylin
// Date		: 2022-01-17

package leetcode

import (
	"fmt"
	"testing"
)

func TestCircularQueue(t *testing.T) {
	queue := Constructor(3)
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)

	full := queue.IsFull()
	fmt.Println(">>> Queue is full:", full)

	// for cur := queue.head; cur != nil; cur = cur.next {
	// 	fmt.Println(cur.value)
	// }

	queue.DeQueue()
	for cur := queue.head; cur != nil; cur = cur.next {
		fmt.Println(cur.value)
	}

	front := queue.Front()
	fmt.Println(">>> Font:", front)

	tail := queue.Rear()
	fmt.Println(">>> Tail:", tail)

	queue.DeQueue()
	queue.DeQueue()
	empty := queue.IsEmpty()
	fmt.Println(">>> Queue is empty:", empty)

}
