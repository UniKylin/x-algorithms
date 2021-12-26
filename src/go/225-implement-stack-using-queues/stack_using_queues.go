// link		: https://leetcode-cn.com/problems/implement-stack-using-queues/
// Author	: Kylin
// Date		: 2021-12-27

package leetcode

type MyStack struct {
	queue []int
}

func MyStackConstructor() MyStack {
	return MyStack{
		queue: []int{},
	}
}

func (s *MyStack) Push(x int) {
	n := len(s.queue)
	s.queue = append(s.queue, x)

	for ; n > 0; n-- {
		s.queue = append(s.queue, s.queue[0])
		s.queue = s.queue[1:]
	}
}

func (s *MyStack) Pop() int {
	val := s.queue[0]
	s.queue = s.queue[1:]
	return val
}

func (s *MyStack) Top() int {
	return s.queue[0]
}

func (s *MyStack) Empty() bool {
	return len(s.queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := MyStackConstructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
