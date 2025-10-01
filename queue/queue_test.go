package queue_test

import (
	"testing"

	"github.com/provincialig/golimitless/queue"
)

func Test_QueueEnqueueDequeue(t *testing.T) {
	q := queue.New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if size := q.Size(); size != 3 {
		t.Errorf("expected size 3, got %d", size)
	}

	val, ok := q.Dequeue()
	if !ok || val != 1 {
		t.Errorf("expected 1, got %v (ok=%v)", val, ok)
	}

	val, ok = q.Dequeue()
	if !ok || val != 2 {
		t.Errorf("expected 2, got %v (ok=%v)", val, ok)
	}

	val, ok = q.Dequeue()
	if !ok || val != 3 {
		t.Errorf("expected 3, got %v (ok=%v)", val, ok)
	}

	if !q.IsEmpty() {
		t.Errorf("expected queue to be empty")
	}
}

func Test_QueueFront(t *testing.T) {
	q := queue.New[string]()
	q.Enqueue("a")
	q.Enqueue("b")

	val, ok := q.Front()
	if !ok || val != "a" {
		t.Errorf("expected front 'a', got %v (ok=%v)", val, ok)
	}

	q.Dequeue()
	val, ok = q.Front()
	if !ok || val != "b" {
		t.Errorf("expected front 'b', got %v (ok=%v)", val, ok)
	}
}

func Test_QueueClear(t *testing.T) {
	q := queue.New[int]()
	q.Enqueue(10)
	q.Enqueue(20)
	q.Clear()

	if !q.IsEmpty() {
		t.Errorf("expected queue to be empty after Clear")
	}

	if size := q.Size(); size != 0 {
		t.Errorf("expected size 0 after Clear, got %d", size)
	}
}

func Test_QueueDequeueEmpty(t *testing.T) {
	q := queue.New[int]()
	val, ok := q.Dequeue()
	if ok {
		t.Errorf("expected ok=false when dequeueing empty queue, got ok=true and val=%v", val)
	}
}

func Test_QueueFrontEmpty(t *testing.T) {
	q := queue.New[int]()
	val, ok := q.Front()
	if ok {
		t.Errorf("expected ok=false when calling Front on empty queue, got ok=true and val=%v", val)
	}
}
