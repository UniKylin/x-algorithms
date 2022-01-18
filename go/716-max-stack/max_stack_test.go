package leetcode

import (
	"testing"
)

func TestMinStack(t *testing.T) {
	stk := MaxStackConstructor()

	stk.Push(5)   // [5] - 5 既是栈顶元素，也是最大元素
	stk.Push(1)   // [5, 1] - 栈顶元素是 1，最大元素是 5
	stk.Push(5)   // [5, 1, 5] - 5 既是栈顶元素，也是最大元素
	stk.Top()     // 返回 5，[5, 1, 5] - 栈没有改变
	stk.PopMax()  // 返回 5，[5, 1] - 栈发生改变，栈顶元素不再是最大元素
	stk.Top()     // 返回 1，[5, 1] - 栈没有改变
	stk.PeekMax() // 返回 5，[5, 1] - 栈没有改变
	stk.Pop()     // 返回 1，[5] - 此操作后，5 既是栈顶元素，也是最大元素
	stk.Top()     // 返回 5，[5] - 栈没有改变

}
