// https://en.wikipedia.org/wiki/Linked_list#Singly_linked_list
package linkedlist

import "fmt"

type Node struct {
	Value interface{}

	// Previous node
	Next *Node
}

type SinglyNode struct {
	Len int

	Head *Node
}

func NewNode(val interface{}) *Node {
	return &Node{val, nil}
}

// Create a new instance of singly linked list node
func NewSinglyNode() *SinglyNode {
	// node := SinglyNode{}
	return &SinglyNode{}
}

func (s *SinglyNode) Add2Head(val interface{}) {
	node := NewNode(val)

	node.Next = s.Head
	s.Head = node
	s.Len++
}

func (s *SinglyNode) Add2Tail(val interface{}) {
	node := NewNode(val)

	if s.Head == nil {
		s.Head = node
		s.Len++
		return
	}

	cur := s.Head

	for ; cur.Next != nil; cur = cur.Next {
		// To the end
	}

	cur.Next = node
	s.Len++
}

func (s *SinglyNode) Length() int {
	return s.Len
}

// Remove node from head and return node's value
// Return -1 if the list is empty
func (s *SinglyNode) RemoveFromHead() interface{} {
	if s.Head == nil {
		return -1
	}

	tempNode := s.Head
	s.Head = tempNode.Next
	s.Len--

	return tempNode.Value
}

func (s *SinglyNode) RemoveFromTail() interface{} {
	// empty list
	if s.Head == nil {
		return -1
	}

	// only one node
	if s.Head.Next == nil {
		val := s.RemoveFromHead()
		return val
	}

	cur := s.Head

	for ; cur.Next.Next != nil; cur = cur.Next {
		// to to to the end, hah...
	}

	val := cur.Next.Value
	cur.Next = nil
	s.Len--

	return val
}

// Reverse singly linked list
func (s *SinglyNode) Reverse() {
	var prev, next *Node

	cur := s.Head

	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	s.Head = prev
}

func (s *SinglyNode) Show() {
	for cur := s.Head; cur != nil; cur = cur.Next {
		if cur.Next != nil {
			fmt.Print(cur.Value, " -> ")
		} else {
			fmt.Print(cur.Value)
		}
	}
	fmt.Println()
}
