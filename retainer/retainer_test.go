package retainer_test

import (
	"sync"
	"testing"
	"time"

	"github.com/provincialig/golimitless/retainer"
)

func TestRetainerAddAndGet(t *testing.T) {
	r := retainer.New[int]()
	defer r.Clean()

	r.Add(42, time.Second)

	ch, cancel := r.Get()
	defer cancel()

	select {
	case v := <-ch:
		if v != 42 {
			t.Errorf("expected 42, got %v", v)
		}
	case <-time.After(2 * time.Second):
		t.Error("timeout waiting for value")
	}
}

func TestRetainerCancel(t *testing.T) {
	r := retainer.New[int]()
	defer r.Clean()

	ch, cancel := r.Get()

	cancel()

	_, ok := <-ch
	if ok {
		t.Error("expected channel to be closed after cancel")
	}
}

func TestRetainerMultipleAdd(t *testing.T) {
	r := retainer.New[string]()
	defer r.Clean()

	r.Add("a", time.Second)
	r.Add("b", time.Second)

	ch, cancel := r.Get()
	defer cancel()

	got := []string{}
	timeout := time.After(3 * time.Second)
	for range 2 {
		select {
		case v := <-ch:
			got = append(got, v)
		case <-timeout:
			t.Fatal("timeout waiting for values")
		}
	}

	if got[0] != "a" || got[1] != "b" {
		t.Errorf("expected [a b], got %v", got)
	}
}

func TestRetainerRetainTiming(t *testing.T) {
	r := retainer.New[string]()
	defer r.Clean()

	r.Add("short", 200*time.Millisecond)
	r.Add("long", 800*time.Millisecond)

	ch, cancel := r.Get()
	defer cancel()

	start := time.Now()
	var first, second string
	var firstTime, secondTime time.Duration

	for i := range 2 {
		select {
		case v := <-ch:
			elapsed := time.Since(start)
			if i == 0 {
				first = v
				firstTime = elapsed
			} else {
				second = v
				secondTime = elapsed
			}
		case <-time.After(2 * time.Second):
			t.Fatal("timeout waiting for values")
		}
	}

	if first != "short" || second != "long" {
		t.Errorf("expected order [short long], got [%s %s]", first, second)
	}
	if firstTime < 180*time.Millisecond || firstTime > 400*time.Millisecond {
		t.Errorf("first retain time out of expected range: %v", firstTime)
	}
	if secondTime < 700*time.Millisecond || secondTime > 1*time.Second {
		t.Errorf("second retain time out of expected range: %v", secondTime)
	}
}

func TestRetainerZeroRetain(t *testing.T) {
	r := retainer.New[int]()
	defer r.Clean()

	r.Add(99, 0)

	ch, cancel := r.Get()
	defer cancel()

	select {
	case v := <-ch:
		if v != 99 {
			t.Errorf("expected 99, got %v", v)
		}
	case <-time.After(200 * time.Millisecond):
		t.Error("timeout waiting for value with zero retain")
	}
}

func TestRetainerNegativeRetain(t *testing.T) {
	r := retainer.New[int]()
	defer r.Clean()

	r.Add(123, -1*time.Second)

	ch, cancel := r.Get()
	defer cancel()

	select {
	case v := <-ch:
		if v != 123 {
			t.Errorf("expected 123, got %v", v)
		}
	case <-time.After(200 * time.Millisecond):
		t.Error("timeout waiting for value with negative retain")
	}
}

func TestRetainerClean(t *testing.T) {
	r := retainer.New[int]()
	r.Add(1, 100*time.Millisecond)
	r.Add(2, 200*time.Millisecond)
	r.Clean()

	ch, cancel := r.Get()
	defer cancel()

	select {
	case v, ok := <-ch:
		if ok {
			t.Errorf("expected channel to be empty after Clean, got %v", v)
		}
	case <-time.After(300 * time.Millisecond):
		// Success: nothing should be received
	}
}

func TestRetainerConcurrentAddAndGet(t *testing.T) {
	r := retainer.New[int]()
	defer r.Clean()

	const n = 50
	done := make(chan struct{})
	ch, cancel := r.Get()
	defer cancel()

	go func() {
		for i := range n {
			r.Add(i, 100*time.Millisecond)
		}
		close(done)
	}()

	received := map[int]bool{}
	timeout := time.After(2 * time.Second)
	for range n {
		select {
		case v := <-ch:
			received[v] = true
		case <-timeout:
			t.Fatal("timeout waiting for concurrent values")
		}
	}
	<-done

	for i := range n {
		if !received[i] {
			t.Errorf("missing value %d in concurrent test", i)
		}
	}
}

func TestRetainerConcurrentCleanAndAdd(t *testing.T) {
	r := retainer.New[int]()
	defer r.Clean()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := range 10 {
			r.Add(i, 100*time.Millisecond)
			time.Sleep(10 * time.Millisecond)
		}
	}()

	go func() {
		defer wg.Done()
		for range 5 {
			r.Clean()
			time.Sleep(20 * time.Millisecond)
		}
	}()

	wg.Wait()
}

func TestRetainerMultipleConsumers(t *testing.T) {
	const (
		VALUES    = 20
		CONSUMERS = 5
	)

	r := retainer.New[int]()

	for i := range VALUES {
		r.Add(i, 100*time.Millisecond)
	}

	var (
		wg      sync.WaitGroup
		mut     sync.Mutex
		results = []int{}
	)

	for range CONSUMERS {
		wg.Add(1)

		go func() {
			defer wg.Done()

			ch, cancel := r.Get()
			defer cancel()

			for v := range ch {
				mut.Lock()
				results = append(results, v)
				mut.Unlock()
			}
		}()
	}

	time.Sleep(500 * time.Millisecond)
	r.Destroy()

	wg.Wait()

	if len(results) != VALUES {
		t.Fail()
	}
}
