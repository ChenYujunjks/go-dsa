// doublylinkedlist_test.go
package linkedlist

import (
	"reflect"
	"testing"
)

func TestDoublyLinkedList_Append(t *testing.T) {
	l := &DoublyLinkedList{}
	l.Append(10)
	l.Append(20)
	l.Append(30)

	expected := []any{10, 20, 30}
	result := ToSlice(l)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Append failed. Expected %v, got %v", expected, result)
	}
}

func TestDoublyLinkedList_Prepend(t *testing.T) {
	l := &DoublyLinkedList{}
	l.Append(20)
	l.Prepend(10)

	expected := []any{10, 20}
	result := ToSlice(l)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Prepend failed. Expected %v, got %v", expected, result)
	}
}

func TestDoublyLinkedList_Delete(t *testing.T) {
	l := &DoublyLinkedList{}
	l.Append(1)
	l.Append(2)
	l.Append(3)

	ok := l.Delete(2)
	if !ok {
		t.Error("Delete should return true when deletion occurs")
	}

	expected := []any{1, 3}
	result := ToSlice(l)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Delete failed. Expected %v, got %v", expected, result)
	}
}

func TestDoublyLinkedList_Find(t *testing.T) {
	l := &DoublyLinkedList{}
	l.Append("a")
	l.Append("b")
	l.Append("c")

	node := l.Find("b")
	if node == nil || node.Value != "b" {
		t.Error("Find failed to locate existing value")
	}

	node = l.Find("x")
	if node != nil {
		t.Error("Find should return nil for non-existent value")
	}
}

func TestDoublyLinkedList_Length(t *testing.T) {
	l := &DoublyLinkedList{}
	l.Append(1)
	l.Append(2)
	l.Append(3)

	if Length(l) != 3 {
		t.Errorf("Length failed. Expected 3, got %d", Length(l))
	}
}

func TestDoublyLinkedList_ToSlice_Empty(t *testing.T) {
	l := &DoublyLinkedList{}
	result := ToSlice(l)
	if len(result) != 0 {
		t.Errorf("ToSlice on empty list failed. Expected [], got %v", result)
	}
}
