package leetcode

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	ant := new(ListNode)
	ant_next := new(ListNode)
	ant_last := new(ListNode)

	ant.Val = 2
	ant_next.Val = 4
	ant_last.Val = 3

	ant.Next = ant_next
	ant_next.Next = ant_last
	ant_last.Next = nil

	node := ant

	for node != nil {
		fmt.Print(node.Val, "-")
		node = node.Next
	}

	fmt.Println()

	bear_last := ListNode{4, nil}
	bear_next := ListNode{6, &bear_last}
	bear := &ListNode{5, &bear_next}

	p := bear

	for p != nil {
		fmt.Print(p.Val, "-")
		p = p.Next
	}

	fmt.Println()

	nums := AddTwoNumbers(ant, bear)
	num := nums
	for num != nil {
		if num.Val >= 0 {
			fmt.Print(num.Val, "-")
		}
		num = num.Next
	}

	fmt.Println()

}
