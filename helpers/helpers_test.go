package helpers_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/provincialig/golimitless/helpers"
)

func Test_MaxRetrySuccess(t *testing.T) {
	counter := 0

	err := helpers.ContinousRetry(context.Background(), 0, 0, func() error {
		if counter == 5 {
			return nil
		}

		counter++
		return errors.New("")
	})

	if err != nil {
		t.Fatal(err)
	}

	if counter != 5 {
		t.Fatal(counter)
	}
}

func Test_MaxRetryFail(t *testing.T) {
	err := helpers.ContinousRetry(context.Background(), 0, 10, func() error {
		return errors.New("")
	})
	if err == nil {
		t.Fail()
	}
}

func Test_ContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	calls := 0

	err := helpers.ContinousRetry(ctx, 10*time.Millisecond, 0, func() error {
		calls++
		return errors.New("fail")
	})
	if err != nil {
		t.Fatalf("expected nil on context cancel, got %v", err)
	}

	if calls == 0 {
		t.Fatalf("expected at least one retry call")
	}
}

func Test_JoinChannels(t *testing.T) {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	for i := 0; i < 100; i++ {
		ch1 <- i
		ch2 <- i * 10
	}

	close(ch1)
	close(ch2)

	out := helpers.JoinChannels([]<-chan int{ch1, ch2})

	count := 0
	for range out {
		count++
	}

	if count != 200 {
		t.Fatalf("expected 200 merged elements, got %d", count)
	}
}
