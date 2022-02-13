// link		: https://leetcode-cn.com/problems/sort-of-stacks-lcci/
// Author	: Kylin
// Date		: 2022-02-13

package leetcode

import (
	"fmt"
	"testing"
)

func TestSortedStack(t *testing.T) {
	s := Constructor()

	s.Push(9)
	s.Pop()

	empty := s.IsEmpty()
	fmt.Println(">>> sorted stack is empty: ", empty)

	s.Push(1)
	s.Push(7)
	s.Push(4)

	fmt.Println(s.data)

	ele := s.Peek()
	fmt.Println(">>> element: ", ele)

}
