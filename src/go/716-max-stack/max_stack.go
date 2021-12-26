// link		: https://leetcode-cn.com/problems/max-stack/
// Author	: Kylin
// Date		: 2021-12-27

package leetcode

import "math"

type MaxStack struct {
	stack    []int
	maxStack []int
}

func MaxStackConstructor() MaxStack {
	return MaxStack{
		stack:    []int{},
		maxStack: []int{math.MinInt64},
	}
}

func (m *MaxStack) Push(x int) {
	m.stack = append(m.stack, x)
	max := x

	if len(m.maxStack) != 0 {
		max = m.PeekMax()
	}
	m.maxStack = append(m.maxStack, int(math.Max(float64(max), float64(x))))
}

func (m *MaxStack) Pop() int {
	val := m.stack[len(m.stack)-1]
	m.stack = m.stack[:len(m.stack)-1]
	m.maxStack = m.maxStack[:len(m.maxStack)-1]
	return val
}

func (m *MaxStack) Top() int {
	return m.stack[len(m.stack)-1]
}

func (m *MaxStack) PeekMax() int {
	return m.maxStack[len(m.maxStack)-1]
}

func (m *MaxStack) PopMax() int {
	max := m.PeekMax()
	var temp []int

	for m.Top() != max {
		temp = append(temp, m.Pop())
	}
	m.Pop()

	for len(temp) != 0 {
		m.Push(temp[len(temp)-1])
		temp = temp[:len(temp)-1]
	}

	return max
}

/**
 * Your MaxStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.PeekMax();
 * param_5 := obj.PopMax();
 */
