// link		: https://leetcode-cn.com/problems/min-stack/
// Author	: Kylin
// Date		: 2021-12-27

package leetcode

type MinStack struct {
	p   []int
	min int
}

func MinStackConstructor() MinStack {
	return MinStack{}
}

func (m *MinStack) Push(val int) {
	if len(m.p) == 0 || m.min > val {
		m.min = val
	}
	m.p = append(m.p, val)
}

func (m *MinStack) Pop() {
	if len(m.p) <= 0 {
		return
	}
	if m.min == m.p[len(m.p)-1] {
		m.min = m.p[0]
		for i := 0; i < len(m.p)-1; i++ {
			if m.p[i] < m.min {
				m.min = m.p[i]
			}
		}
	}
	m.p = m.p[:len(m.p)-1]
}

func (m *MinStack) Top() int {
	if len(m.p) <= 0 {
		return 0
	}
	return m.p[len(m.p)-1]
}

func (m *MinStack) GetMin() int {
	return m.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
