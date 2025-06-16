package stack

import (
	"testing"

	l "github.com/ChenYujunjks/go-review/ds/link"
	"github.com/ChenYujunjks/go-review/ds/utils"
)

func TestStack_PushPopPeek(t *testing.T) {
	s := &Stack{list: &l.LinkedList{}}

	s.Push("a")
	s.Push("b")
	s.Push("c")

	// 栈顶应该是 c
	val, ok := s.Peek()
	if !ok || val != "c" {
		t.Errorf("Peek failed. Expected 'c', got %v", val)
	}

	if s.Len() != 3 {
		t.Errorf("Expected length 3, got %d", s.Len())
	}

	// Pop 应该依次得到 c, b, a
	expected := []string{"c", "b", "a"}
	for i, exp := range expected {
		val, ok := s.Pop()
		if !ok || val != exp {
			t.Errorf("Pop %d failed. Expected %v, got %v", i, exp, val)
		}
	}

	// 此时应为空
	if !s.IsEmpty() {
		t.Error("Expected stack to be empty after popping all elements")
	}
}

func TestStack_EmptyPop(t *testing.T) {
	s := &Stack{list: &l.LinkedList{}}

	val, ok := s.Pop()
	if ok || val != nil {
		t.Errorf("Pop on empty stack should return nil, false. Got %v, %v", val, ok)
	}
}

func TestStack_EmptyPeek(t *testing.T) {
	s := &Stack{list: &l.LinkedList{}}

	val, ok := s.Peek()
	if ok || val != nil {
		t.Errorf("Peek on empty stack should return nil, false. Got %v, %v", val, ok)
	}
}

func TestStack_LenAndIsEmpty(t *testing.T) {
	s := &Stack{list: &l.LinkedList{}}

	if !s.IsEmpty() {
		t.Error("Expected new stack to be empty")
	}

	s.Push(1)
	s.Push(2)

	if s.Len() != 2 {
		t.Errorf("Expected length 2, got %d", s.Len())
	}

	s.Pop()
	s.Pop()

	if !s.IsEmpty() {
		t.Error("Expected stack to be empty after popping all elements")
	}
}

func TestStack_ImplementsLL(t *testing.T) {
	s := New()

	s.Push(10)
	s.Push(20)
	s.Push(30)

	slice := utils.ToSlice(s)

	expected := []any{30, 20, 10}
	for i, val := range expected {
		if slice[i] != val {
			t.Errorf("Expected %v at index %d, got %v", val, i, slice[i])
		}
	}
}
