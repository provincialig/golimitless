package retrier

import (
	"context"
	"errors"
	"time"
)

const (
	NO_DELAY     time.Duration = 0
	NO_MAX_RETRY int           = 0
)

var (
	ErrMaxRetry       = errors.New("max retry")
	ErrContextCancel  = errors.New("context cancel")
	ErrContextTimeout = errors.New("context timeout")
)

type Retrier interface {
	Run(ctx context.Context, fn func() error) error
}

type myRetrier struct {
	delay    time.Duration
	maxRetry int
}

func (r myRetrier) Run(ctx context.Context, fn func() error) error {
	retry := 0

	for {
		if ctx.Err() != nil {
			if errors.Is(ctx.Err(), context.Canceled) {
				return ErrContextCancel
			}
			if errors.Is(ctx.Err(), context.DeadlineExceeded) {
				return ErrContextTimeout
			}
			return ctx.Err()
		}

		if r.maxRetry > 0 && retry >= r.maxRetry {
			return ErrMaxRetry
		}

		retry++

		if err := fn(); err != nil {
			if r.delay == 0 {
				continue
			}

			select {
			case <-time.After(r.delay):
			case <-ctx.Done():
				if errors.Is(ctx.Err(), context.Canceled) {
					return ErrContextCancel
				}
				if errors.Is(ctx.Err(), context.DeadlineExceeded) {
					return ErrContextTimeout
				}
				return ctx.Err()
			}
		} else {
			return nil
		}
	}
}

func New(delay time.Duration, maxRetry int) Retrier {
	return myRetrier{
		delay:    delay,
		maxRetry: maxRetry,
	}
}
