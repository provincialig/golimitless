package stack_test

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/provincialig/golimitless/stack"
)

func Test_Push(t *testing.T) {
	s := stack.New[int]()
	s.Push(10)
	s.Push(20)

	size := s.Size()
	if size != 2 {
		t.Fatalf("Size: %d", size)
	}

	if s.IsEmpty() {
		t.Fatal("Stack must be enmpty")
	}
}

func Test_Pop(t *testing.T) {
	s := stack.New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	if el, ok := s.TryPop(); !ok || el != 3 {
		t.Fatalf("Element: %d", el)
	}

	if size := s.Size(); size != 2 {
		t.Fatalf("Size: %d", size)
	}

	if el, ok := s.TryPop(); !ok || el != 2 {
		t.Fatalf("Element: %d", el)
	}

	if size := s.Size(); size != 1 {
		t.Fatalf("Size: %d", size)
	}

	if el, ok := s.TryPop(); !ok || el != 1 {
		t.Fatalf("Element: %d", el)
	}

	if size := s.Size(); size != 0 {
		t.Fatalf("Size: %d", size)
	}

	if !s.IsEmpty() {
		t.Fatal("Is empty", s.IsEmpty())
	}
}

func Test_Peek(t *testing.T) {
	s := stack.New[int]()
	s.Push(10)
	s.Push(20)

	if el, ok := s.TryPeek(); !ok || el != 20 {
		t.Fatalf("Element: %d", el)
	}

	if size := s.Size(); size != 2 {
		t.Fatalf("Size: %d", size)
	}
}

func Test_Clear(t *testing.T) {
	s := stack.New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	initialSize := s.Size()
	if initialSize != 3 {
		t.Fatalf("Initial size: %d", initialSize)
	}

	s.Clear()

	finalSize := s.Size()
	if finalSize != 0 {
		t.Fatalf("Final size: %d", initialSize)
	}
}

func Test_Pop_Blocking_Success_Multithread(t *testing.T) {
	s := stack.New[int]()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
		s.Push(42)
	}()
	val, err := s.Pop(ctx)
	wg.Wait()
	if err != nil {
		t.Fatalf("Pop error: %v", err)
	}
	if val != 42 {
		t.Fatalf("Expected 42, got %d", val)
	}
}

func Test_Pop_Blocking_Timeout_Multithread(t *testing.T) {
	s := stack.New[int]()
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(200 * time.Millisecond)
		s.Push(99)
	}()
	val, err := s.Pop(ctx)
	wg.Wait()
	if !errors.Is(err, stack.ErrTimeout) {
		t.Fatalf("Expected ErrTimeout, got %v", err)
	}
	if val != 0 {
		t.Fatalf("Expected zero value, got %d", val)
	}
}

func Test_Pop_Blocking_Canceled_Multithread(t *testing.T) {
	s := stack.New[int]()
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()
	val, err := s.Pop(ctx)
	wg.Wait()
	if !errors.Is(err, stack.ErrCanceled) {
		t.Fatalf("Expected ErrCanceled, got %v", err)
	}
	if val != 0 {
		t.Fatalf("Expected zero value, got %d", val)
	}
}

func Test_Peek_Blocking_Success_Multithread(t *testing.T) {
	s := stack.New[int]()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
		s.Push(99)
	}()
	val, err := s.Peek(ctx)
	wg.Wait()
	if err != nil {
		t.Fatalf("Peek error: %v", err)
	}
	if val != 99 {
		t.Fatalf("Expected 99, got %d", val)
	}
}

func Test_Peek_Blocking_Timeout_Multithread(t *testing.T) {
	s := stack.New[int]()
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(200 * time.Millisecond)
		s.Push(88)
	}()
	val, err := s.Peek(ctx)
	wg.Wait()
	if !errors.Is(err, stack.ErrTimeout) {
		t.Fatalf("Expected ErrTimeout, got %v", err)
	}
	if val != 0 {
		t.Fatalf("Expected zero value, got %d", val)
	}
}

func Test_Peek_Blocking_Canceled_Multithread(t *testing.T) {
	s := stack.New[int]()
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()
	val, err := s.Peek(ctx)
	wg.Wait()
	if !errors.Is(err, stack.ErrCanceled) {
		t.Fatalf("Expected ErrCanceled, got %v", err)
	}
	if val != 0 {
		t.Fatalf("Expected zero value, got %d", val)
	}
}

func Test_TryPop_Empty_Multithread(t *testing.T) {
	s := stack.New[int]()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Millisecond)
	}()
	val, ok := s.TryPop()
	wg.Wait()
	if ok {
		t.Fatalf("Expected not ok, got ok with value %d", val)
	}
}

func Test_TryPeek_Empty_Multithread(t *testing.T) {
	s := stack.New[int]()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Millisecond)
	}()
	val, ok := s.TryPeek()
	wg.Wait()
	if ok {
		t.Fatalf("Expected not ok, got ok with value %d", val)
	}
}
