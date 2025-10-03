package queue_test

import (
	"context"
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

	if val, err := q.Dequeue(context.Background()); err != nil || val != 1 {
		t.Errorf("expected 1, got %v %v", val, err)
	}

	if val, err := q.Dequeue(context.Background()); err != nil || val != 2 {
		t.Errorf("expected 2, got %v %v", val, err)
	}

	if val, err := q.Dequeue(context.Background()); err != nil || val != 3 {
		t.Errorf("expected 3, got %v %v", val, err)
	}

	if !q.IsEmpty() {
		t.Errorf("expected queue to be empty")
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

func Test_QueueTryDequeue(t *testing.T) {
	q := queue.New[int]()

	// TryDequeue on empty queue
	val, ok := q.TryDequeue()
	if ok {
		t.Errorf("expected ok=false on empty queue, got ok=true with value %v", val)
	}

	// Enqueue some elements
	q.Enqueue(100)
	q.Enqueue(200)

	// TryDequeue should return first element
	val, ok = q.TryDequeue()
	if !ok || val != 100 {
		t.Errorf("expected ok=true and value=100, got ok=%v value=%v", ok, val)
	}

	// TryDequeue should return second element
	val, ok = q.TryDequeue()
	if !ok || val != 200 {
		t.Errorf("expected ok=true and value=200, got ok=%v value=%v", ok, val)
	}

	// Queue should be empty now
	val, ok = q.TryDequeue()
	if ok {
		t.Errorf("expected ok=false after all elements dequeued, got ok=true with value %v", val)
	}
}
