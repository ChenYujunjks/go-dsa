package list_test

import (
	"reflect"
	"testing"

	list "github.com/ChenYujunjks/go-review/ds/link/list"
)

func TestList_Append(t *testing.T) {
	l := list.NewList()
	l.Append(10)
	l.Append(20)
	l.Append(30)

	expected := []any{10, 20, 30}
	result := l.ToSlice()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Append failed. Expected %v, got %v", expected, result)
	}
}

func TestList_Prepend(t *testing.T) {
	l := list.NewList()
	l.Append(20)
	l.Prepend(10)

	expected := []any{10, 20}
	result := l.ToSlice()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Prepend failed. Expected %v, got %v", expected, result)
	}
}

func TestList_Insert(t *testing.T) {
	l := list.NewList()
	l.Append(1)
	l.Append(3)
	l.Insert(1, 2) // 插入中间

	expected := []any{1, 2, 3}
	result := l.ToSlice()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Insert failed. Expected %v, got %v", expected, result)
	}
}

func TestList_Delete(t *testing.T) {
	l := list.NewList()
	l.Append(1)
	l.Append(2)
	l.Append(3)

	ok := l.Delete(2)
	if !ok {
		t.Error("Delete should return true when deletion occurs")
	}

	expected := []any{1, 3}
	result := l.ToSlice()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Delete failed. Expected %v, got %v", expected, result)
	}
}

func TestList_Find(t *testing.T) {
	l := list.NewList()
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

func TestList_Length(t *testing.T) {
	l := list.NewList()
	l.Append(1)
	l.Append(2)
	l.Append(3)

	if l.Length() != 3 {
		t.Errorf("Length failed. Expected 3, got %d", l.Length())
	}
}

func TestList_ToSlice_Empty(t *testing.T) {
	l := list.NewList()
	result := l.ToSlice()
	if len(result) != 0 {
		t.Errorf("ToSlice on empty list failed. Expected [], got %v", result)
	}
}
