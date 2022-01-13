package linkedlist

import (
	"fmt"
	"testing"
)

func TestSinglyLinkedList(t *testing.T) {
	list := NewSinglyNode()
	list.Add2Head(1)
	list.Add2Head(2)
	list.Add2Head(3)
	list.Add2Head(4)
	list.Show()

	list.Add2Tail(6)
	list.Add2Tail(7)
	list.Add2Tail(8)
	list.Show()

	headVal := list.RemoveFromHead()
	fmt.Println(">>> Head value:", headVal)

	tailVal := list.RemoveFromTail()
	fmt.Println(">>> Tail value:", tailVal)

	list.Show()

	length := list.Length()
	fmt.Println(">>> Singly Linked List Length:", length)

}

func BenchmarkSinglyLinkedList(b *testing.B) {
	for n := 0; n < b.N; n++ {
		list := NewSinglyNode()
		list.Add2Head(1)
		list.Add2Head(2)
		list.Add2Head(3)
		list.Add2Head(4)

		list.Add2Tail(6)
		list.Add2Tail(7)
		list.Add2Tail(8)

		list.RemoveFromHead()
		list.RemoveFromTail()

		list.Length()
	}
}
