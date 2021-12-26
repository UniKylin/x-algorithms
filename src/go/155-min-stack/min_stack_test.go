package leetcode

import (
	"fmt"
	"testing"
)

func TestMinStack(t *testing.T) {
	minStack := MinStackConstructor()

	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(">>>", minStack.GetMin()) // --> return -3.
	minStack.Pop()
	fmt.Println(">>>", minStack.Top())    //--> return 0.
	fmt.Println(">>>", minStack.GetMin()) //--> return -2.

}
