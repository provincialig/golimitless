package helpers_test

import (
	"testing"

	"github.com/provincialig/golimitless/helpers"
)

func Test_JoinChannels(t *testing.T) {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	for i := range 100 {
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
