package leetcode

import (
	"fmt"
	"testing"
)

func TestQueueUsingStack(t *testing.T) {
	myQueue := MyQueueConstructor()

	myQueue.Push(1)                    // queue is: [1]
	myQueue.Push(2)                    // queue is: [1, 2] (leftmost is front of the queue)
	myQueue.Push(3)                    // queue is: [1, 2] (leftmost is front of the queue)
	fmt.Println(">>>", myQueue.Peek()) // return 1
	fmt.Println(">>>", myQueue.Pop())  // return 1, queue is [2]
	myQueue.Empty()                    // return false

}
