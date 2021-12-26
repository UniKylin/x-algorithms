package leetcode

import (
	"fmt"
	"testing"
)

func TestMyStack(t *testing.T) {
	myStack := MyStackConstructor()

	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)

	p1 := myStack.Pop()
	fmt.Println(">>> p1:", p1)

	t1 := myStack.Top()
	fmt.Println(">>> top:", t1)

	isEmpty := myStack.Empty()
	fmt.Println(">>> stack is empty:", isEmpty)

}
