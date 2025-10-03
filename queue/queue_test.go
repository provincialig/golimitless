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

	if val := q.Dequeue(); val != 1 {
		t.Errorf("expected 1, got %v", val)
	}

	if val := q.Dequeue(); val != 2 {
		t.Errorf("expected 2, got %v", val)
	}

	if val := q.Dequeue(); val != 3 {
		t.Errorf("expected 3, got %v", val)
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
