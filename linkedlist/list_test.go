package linkedlist_test

import (
	"slices"
	"testing"

	"github.com/Kidpech-code/algorithms-made-clear/linkedlist"
)

func TestListOperations(t *testing.T) {
	var head *linkedlist.Node
	for _, value := range []int{5, 1, 4, 2, 8, 3} {
		head = linkedlist.Append(head, value)
	}
	wantValues(t, head, []int{5, 1, 4, 2, 8, 3})

	if got := linkedlist.IndexOf(head, 2); got != 3 {
		t.Fatalf("IndexOf(2) = %d; want 3", got)
	}
	if got := linkedlist.IndexOf(head, 9); got != -1 {
		t.Fatalf("IndexOf(9) = %d; want -1", got)
	}

	head = linkedlist.Prepend(head, 9)
	wantValues(t, head, []int{9, 5, 1, 4, 2, 8, 3})
	head = linkedlist.DeleteFirst(head, 4)
	wantValues(t, head, []int{9, 5, 1, 2, 8, 3})
	head = linkedlist.DeleteFirst(head, 9)
	wantValues(t, head, []int{5, 1, 2, 8, 3})
	head = linkedlist.DeleteFirst(head, 3)
	wantValues(t, head, []int{5, 1, 2, 8})
	head = linkedlist.DeleteFirst(head, 99)
	wantValues(t, head, []int{5, 1, 2, 8})
}

func TestEmptyAndDuplicateValues(t *testing.T) {
	var head *linkedlist.Node
	if got := linkedlist.IndexOf(head, 4); got != -1 {
		t.Fatalf("IndexOf on empty list = %d; want -1", got)
	}
	if got := linkedlist.DeleteFirst(head, 4); got != nil {
		t.Fatal("DeleteFirst on empty list should return nil")
	}

	head = linkedlist.Prepend(head, 4)
	head = linkedlist.Append(head, 2)
	head = linkedlist.Append(head, 4)
	if got := linkedlist.IndexOf(head, 4); got != 0 {
		t.Fatalf("IndexOf first duplicate = %d; want 0", got)
	}
	head = linkedlist.DeleteFirst(head, 4)
	wantValues(t, head, []int{2, 4})
}

func wantValues(t *testing.T, head *linkedlist.Node, want []int) {
	t.Helper()
	var got []int
	for node := head; node != nil && len(got) <= len(want); node = node.Next {
		got = append(got, node.Value)
	}
	if !slices.Equal(got, want) {
		t.Fatalf("list values = %v; want %v", got, want)
	}
}
