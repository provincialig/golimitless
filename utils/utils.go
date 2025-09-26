package utils

import (
	"context"
	"sync"
	"time"
)

func JoinChannels[T any](subscriptions []<-chan T, bufferSize ...int) <-chan T {
	buff := 1000
	if len(bufferSize) > 0 {
		buff = bufferSize[0]
	}

	var wg sync.WaitGroup
	out := make(chan T, buff)

	go func() {
		for _, ch := range subscriptions {
			wg.Add(1)
			go func(c <-chan T, wg *sync.WaitGroup) {
				for val := range c {
					out <- val
				}
				wg.Done()
			}(ch, &wg)
		}

		wg.Wait()
		close(out)
	}()

	return out
}

func ContinousRetry(ctx context.Context, sleep time.Duration, fn func() error) {
	for {
		if ctx.Err() != nil {
			return
		}

		if err := fn(); err != nil {
			select {
			case <-time.After(sleep):
			case <-ctx.Done():
				return
			}
		} else {
			return
		}
	}
}
