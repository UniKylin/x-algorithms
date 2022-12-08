// link		: https://leetcode-cn.com/problems/implement-stack-using-queues/
// Author	: Kylin
// Date		: 2021-12-27

package leetcode

type MyQueue struct {
	inStack  []int
	outStack []int
}

func MyQueueConstructor() MyQueue {
	return MyQueue{
		inStack:  []int{},
		outStack: []int{},
	}
}

func (s *MyQueue) Push(x int) {
	s.inStack = append(s.inStack, x)
}

func (s *MyQueue) in2out() {
	for len(s.inStack) > 0 {
		s.outStack = append(s.outStack, s.inStack[len(s.inStack)-1])
		s.inStack = s.inStack[:len(s.inStack)-1]
	}
}

func (s *MyQueue) Pop() int {
	if len(s.outStack) == 0 {
		s.in2out()
	}
	val := s.outStack[len(s.outStack)-1]
	s.outStack = s.outStack[:len(s.outStack)-1]
	return val
}

func (s *MyQueue) Peek() int {
	if len(s.outStack) == 0 {
		s.in2out()
	}

	return s.outStack[len(s.outStack)-1]
}

func (s *MyQueue) Empty() bool {
	return len(s.outStack) == 0 && len(s.inStack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
