// linkedlist_test.go
package linkedlist

import (
	"reflect"
	"testing"
)

func TestLinkedList_Append(t *testing.T) {
	l := &LinkedList{}
	l.Append(10)
	l.Append(20)
	l.Append(30)

	expected := []any{10, 20, 30}
	result := ToSlice(l)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Append failed. Expected %v, got %v", expected, result)
	}
}

func TestLinkedList_Prepend(t *testing.T) {
	l := &LinkedList{}
	l.Append(20)
	l.Prepend(10)

	expected := []any{10, 20}
	result := ToSlice(l)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Prepend failed. Expected %v, got %v", expected, result)
	}
}

func TestLinkedList_Insert(t *testing.T) {
	l := &LinkedList{}
	l.Append(1)
	l.Append(3)
	l.Insert(1, 2) // 插入到中间

	expected := []any{1, 2, 3}
	result := ToSlice(l)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Insert failed. Expected %v, got %v", expected, result)
	}
}

func TestLinkedList_Delete(t *testing.T) {
	l := &LinkedList{}
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

func TestLinkedList_Find(t *testing.T) {
	l := &LinkedList{}
	l.Append("apple")
	l.Append("banana")

	node := l.Find("banana")
	if node == nil || node.Value != "banana" {
		t.Error("Find failed to locate existing value")
	}

	node = l.Find("orange")
	if node != nil {
		t.Error("Find should return nil for non-existent value")
	}
}

func TestLinkedList_Length(t *testing.T) {
	l := &LinkedList{}
	l.Append(1)
	l.Append(2)
	l.Append(3)

	if Length(l) != 3 {
		t.Errorf("Length failed. Expected 3, got %d", Length(l))
	}
}

func TestLinkedList_ToSlice_Empty(t *testing.T) {
	l := &LinkedList{}
	result := ToSlice(l)
	if len(result) != 0 {
		t.Errorf("ToSlice on empty list failed. Expected [], got %v", result)
	}
}
