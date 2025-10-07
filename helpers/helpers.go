package helpers

import (
	"sync"
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
