package queue

import (
	"testing"
)

func TestQueue_EnqueueDequeue(t *testing.T) {
	q := New()
	q.Enqueue("a")
	q.Enqueue("b")
	q.Enqueue("c")

	val, ok := q.Dequeue()
	if !ok || val != "a" {
		t.Errorf("Expected 'a', got %v", val)
	}

	val, ok = q.Dequeue()
	if !ok || val != "b" {
		t.Errorf("Expected 'b', got %v", val)
	}

	val, ok = q.Peek()
	if !ok || val != "c" {
		t.Errorf("Expected peek to be 'c', got %v", val)
	}

	if q.Len() != 1 {
		t.Errorf("Expected length to be 1, got %d", q.Len())
	}

	val, ok = q.Dequeue()
	if !ok || val != "c" {
		t.Errorf("Expected 'c', got %v", val)
	}

	if !q.IsEmpty() {
		t.Error("Expected queue to be empty")
	}
}

func TestQueue_Dequeue_Empty(t *testing.T) {
	q := New()

	val, ok := q.Dequeue()
	if ok || val != nil {
		t.Errorf("Expected dequeue from empty queue to return nil, got %v", val)
	}
}

func TestQueue_Peek_Empty(t *testing.T) {
	q := New()

	val, ok := q.Peek()
	if ok || val != nil {
		t.Errorf("Expected peek from empty queue to return nil, got %v", val)
	}
}
