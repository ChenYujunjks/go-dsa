package queue

import (
	"testing"

	"github.com/ChenYujunjks/go-review/ds/iface"
	"github.com/ChenYujunjks/go-review/ds/utils"
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

// 测试 GetHead 是否和Peek 一样
func TestQueue_GetHead(t *testing.T) {
	q := New()
	q.Enqueue(1)
	q.Enqueue(2)

	head := q.GetHead()
	if head == nil || head.GetValue() != 1 {
		t.Errorf("Expected head value to be 1, got %v", head)
	}
}

// 边界测试 交替出入操作
func TestQueue_Alternating(t *testing.T) {
	q := New()

	q.Enqueue(10)
	val, _ := q.Dequeue()
	if val != 10 {
		t.Errorf("Expected 10, got %v", val)
	}

	q.Enqueue(20)
	q.Enqueue(30)
	val, _ = q.Dequeue()
	if val != 20 {
		t.Errorf("Expected 20, got %v", val)
	}
}

// 测试 Queue 是否实现了 LL 接口
func TestQueue_ImplementsLL(t *testing.T) {
	q := New()

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	// 使用接口类型接收
	var ll iface.LL = q

	// 用 utils.ToSlice 处理接口
	slice := utils.ToSlice(ll)

	expected := []any{10, 20, 30}
	for i, val := range expected {
		if slice[i] != val {
			t.Errorf("Expected %v at index %d, got %v", val, i, slice[i])
		}
	}
}

// 暴力测试 
func TestQueue_Massive(t *testing.T) {
	q := New()

	for i := 0; i < 10000; i++ {
		q.Enqueue(i)
	}
	if q.Len() != 10000 {
		t.Errorf("Expected length 10000, got %d", q.Len())
	}

	for i := 0; i < 10000; i++ {
		val, ok := q.Dequeue()
		if !ok || val != i {
			t.Errorf("Mismatch at %d: got %v", i, val)
		}
	}
	if !q.IsEmpty() {
		t.Errorf("Queue should be empty after all dequeues")
	}
}
