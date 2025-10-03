package queue_test

import (
	"context"
	"sync"
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

func Test_QueueDequeue_BlockingAndCancel(t *testing.T) {
	q := queue.New[int]()
	ctx, cancel := context.WithCancel(context.Background())

	done := make(chan struct{})
	go func() {
		val, err := q.Dequeue(ctx)
		if err != queue.ErrCanceled {
			t.Errorf("expected ErrCanceled, got %v", err)
		}
		if val != 0 {
			t.Errorf("expected zero value, got %v", val)
		}
		close(done)
	}()

	// Cancel after short delay
	cancel()
	<-done
}

func Test_QueueDequeue_BlockingAndTimeout(t *testing.T) {
	q := queue.New[int]()
	ctx, cancel := context.WithTimeout(context.Background(), 0)
	defer cancel()

	val, err := q.Dequeue(ctx)
	if err != queue.ErrTimeout {
		t.Errorf("expected ErrTimeout, got %v", err)
	}
	if val != 0 {
		t.Errorf("expected zero value, got %v", val)
	}
}

func Test_QueueDequeue_ConcurrentEnqueueDequeue(t *testing.T) {
	q := queue.New[int]()
	const n = 100
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Enqueue in goroutine
	go func() {
		for i := range n {
			q.Enqueue(i)
		}
		wg.Done()
	}()

	// Dequeue in goroutine
	go func() {
		for i := range n {
			val, err := q.Dequeue(context.Background())
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if val != i {
				t.Errorf("expected %d, got %d", i, val)
			}
		}
		wg.Done()
	}()

	wg.Wait()
	if !q.IsEmpty() {
		t.Errorf("expected queue to be empty after concurrent ops")
	}
}

func Test_QueueTryDequeue_Concurrent(t *testing.T) {
	q := queue.New[int]()
	const n = 50
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := range n {
			q.Enqueue(i)
		}
		wg.Done()
	}()

	go func() {
		count := 0
		for count < n {
			val, ok := q.TryDequeue()
			if ok {
				count++
				if val < 0 || val >= n {
					t.Errorf("dequeued value out of range: %d", val)
				}
			}
		}
		wg.Done()
	}()

	wg.Wait()
	if !q.IsEmpty() {
		t.Errorf("expected queue to be empty after concurrent TryDequeue")
	}
}
