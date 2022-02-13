// link		: https://leetcode-cn.com/problems/sort-of-stacks-lcci/
// Author	: Kylin
// Date		: 2022-02-13

package leetcode

type SortedStack struct {
	data []int
}

func Constructor() SortedStack {
	return SortedStack{
		data: make([]int, 0),
	}
}

func (s *SortedStack) Push(val int) {
	index := len(s.data)

	if index == 0 || s.data[index-1] <= val {
		s.data = append(s.data, val)
		return
	}

	for index != 0 && s.data[index-1] > val {
		index--
	}

	s.data = append(s.data[:index], append([]int{val}, s.data[index:]...)...)
}

func (s *SortedStack) Pop() {
	if s.IsEmpty() {
		return
	}

	s.data = s.data[1:]
}

func (s *SortedStack) Peek() int {
	if s.IsEmpty() {
		return -1
	}
	return s.data[0]
}

func (s *SortedStack) IsEmpty() bool {
	return len(s.data) == 0
}
