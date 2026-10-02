// Package linkedlist contains a small singly linked list for learning.
package linkedlist

// Node holds one value and a link to the next node. A nil Next ends the list.
type Node struct {
	Value int
	Next  *Node
}

// Prepend adds value before head and returns the new head.
func Prepend(head *Node, value int) *Node {
	return &Node{Value: value, Next: head}
}

// Append adds value after the last node and returns the head. Without a tail
// pointer, it must walk through the list to find the last node.
func Append(head *Node, value int) *Node {
	if head == nil {
		return &Node{Value: value}
	}
	last := head
	for last.Next != nil {
		last = last.Next
	}
	last.Next = &Node{Value: value}
	return head
}

// IndexOf returns the first position containing value, or -1 if absent.
func IndexOf(head *Node, value int) int {
	index := 0
	for node := head; node != nil; node = node.Next {
		if node.Value == value {
			return index
		}
		index++
	}
	return -1
}

// DeleteFirst removes the first node containing value and returns the head,
// which may change when the first node is removed.
func DeleteFirst(head *Node, value int) *Node {
	if head == nil {
		return nil
	}
	if head.Value == value {
		return head.Next
	}
	for node := head; node.Next != nil; node = node.Next {
		if node.Next.Value == value {
			node.Next = node.Next.Next
			return head
		}
	}
	return head
}
