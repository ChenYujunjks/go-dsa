package list

import (
	"reflect"
	"testing"

	"github.com/ChenYujunjks/go-review/ds/utils"
)

func TestList_Append(t *testing.T) {
	l := NewList()
	l.Append(10)
	l.Append(20)
	l.Append(30)

	expected := []any{10, 20, 30}
	result := utils.ToSlice(l)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Append failed. Expected %v, got %v", expected, result)
	}
}

func TestList_Prepend(t *testing.T) {
	l := NewList()
	l.Append(20)
	l.Prepend(10)

	expected := []any{10, 20}
	result := utils.ToSlice(l)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Prepend failed. Expected %v, got %v", expected, result)
	}
}

func TestList_Insert(t *testing.T) {
	l := NewList()
	l.Append(1)
	l.Append(3)
	l.Insert(1, 2) // 插入中间

	expected := []any{1, 2, 3}
	result := utils.ToSlice(l)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Insert failed. Expected %v, got %v", expected, result)
	}
}

func TestList_Delete(t *testing.T) {
	l := NewList()
	l.Append(1)
	l.Append(2)
	l.Append(3)

	ok := l.Delete(2)
	if !ok {
		t.Error("Delete should return true when deletion occurs")
	}

	expected := []any{1, 3}
	result := utils.ToSlice(l)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Delete failed. Expected %v, got %v", expected, result)
	}
}

func TestList_Find(t *testing.T) {
	l := NewList()
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
	l := NewList()
	l.Append(1)
	l.Append(2)
	l.Append(3)

	if utils.Length(l) != 3 {
		t.Errorf("Length failed. Expected 3, got %d", utils.Length(l))
	}
}

func TestList_ToSlice_Empty(t *testing.T) {
	l := NewList()
	result := utils.ToSlice(l)
	if len(result) != 0 {
		t.Errorf("ToSlice on empty list failed. Expected [], got %v", result)
	}
}
func TestList_Get(t *testing.T) {
	l := NewList()
	l.Append("a")
	l.Append("b")
	l.Append("c")

	tests := []struct {
		index    int
		expected interface{}
		ok       bool
	}{
		{0, "a", true},
		{1, "b", true},
		{2, "c", true},
		{3, nil, false},  // 超出范围
		{-1, nil, false}, // 负数下标
	}

	for _, tt := range tests {
		val, ok := l.Get(tt.index)
		if ok != tt.ok || val != tt.expected {
			t.Errorf("Get(%d) = (%v, %v), expected (%v, %v)", tt.index, val, ok, tt.expected, tt.ok)
		}
	}
}
