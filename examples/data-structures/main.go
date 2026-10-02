package main

import (
	"fmt"
	"slices"

	"github.com/Kidpech-code/algorithms-made-clear/linkedlist"
)

func main() {
	values := [6]int{5, 1, 4, 2, 8, 3}
	fmt.Println("Array:", values)
	fmt.Println("Array[3]:", values[3])
	copied := values
	copied[3] = 7
	fmt.Println("Original after changing copy:", values)
	fmt.Println("Copy after update:", copied)
	inserted := slices.Insert(values[:], 2, 7)
	fmt.Println("Slice after insert 7 at index 2:", inserted)

	var head *linkedlist.Node
	for i := len(values) - 1; i >= 0; i-- {
		head = linkedlist.Prepend(head, values[i])
	}
	printList("Linked list: ", head)
	fmt.Println("Index of 2:", linkedlist.IndexOf(head, 2))
	head = linkedlist.Prepend(head, 9)
	printList("After prepend 9: ", head)
	head = linkedlist.DeleteFirst(head, 4)
	printList("After delete 4: ", head)
}

func printList(label string, head *linkedlist.Node) {
	fmt.Print(label)
	for node := head; node != nil; node = node.Next {
		fmt.Printf("%d -> ", node.Value)
	}
	fmt.Println("nil")
}
