package utils

import (
	"context"
	"errors"
	"sync"
	"time"
)

// JoinChannels unisce più channel in un unico.
// Il channel di ritorno verrà chiuso in automatico quando tutti i channel iniziali saranno chiusi.
func JoinChannels[T any](channels []<-chan T, size ...int) <-chan T {
	var (
		wg  sync.WaitGroup
		out chan T
	)

	if len(size) > 0 {
		out = make(chan T, size[0])
	} else {
		out = make(chan T)
	}

	go func() {
		for _, ch := range channels {
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

var ErrMaxRetry = errors.New("max retry")

// ContinuousRetry esegue fn fino al successo (non ritorna più errore) o alla cancellazione del context.
//
// sleep: durata di attesa tra i tentativi (0 = nessuna attesa)
//
// maxRetry: numero massimo di tentativi (0 = illimitato)
func ContinousRetry(ctx context.Context, sleep time.Duration, maxRetry int, fn func() error) error {
	retry := 0

	for {
		if ctx.Err() != nil {
			return nil
		}

		if maxRetry > 0 && retry >= maxRetry {
			return ErrMaxRetry
		}

		retry++

		if err := fn(); err != nil {
			if sleep == 0 {
				continue
			}

			select {
			case <-time.After(sleep):
			case <-ctx.Done():
				return nil
			}
		} else {
			return nil
		}
	}
}
