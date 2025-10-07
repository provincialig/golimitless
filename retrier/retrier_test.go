package retrier_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/provincialig/golimitless/retrier"
)

func Test_MaxRetrySuccess(t *testing.T) {
	counter := 0

	err := retrier.New(retrier.NO_DELAY, retrier.NO_MAX_RETRY).Run(context.Background(), func() error {
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
	err := retrier.New(retrier.NO_DELAY, 10).Run(context.Background(), func() error {
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

	err := retrier.New(10*time.Millisecond, retrier.NO_MAX_RETRY).Run(ctx, func() error {
		calls++
		return errors.New("fail")
	})
	if !errors.Is(err, retrier.ErrContextCancel) {
		t.Fatalf("expected ErrContextCancel on context cancel, got %v", err)
	}

	if calls == 0 {
		t.Fatalf("expected at least one retry call")
	}
}

func Test_ContextTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	err := retrier.New(10*time.Millisecond, retrier.NO_MAX_RETRY).Run(ctx, func() error {
		return errors.New("fail")
	})
	if !errors.Is(err, retrier.ErrContextTimeout) {
		t.Fatalf("expected ErrContextTimeout, got %v", err)
	}
}
