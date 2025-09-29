package utils_test

import (
	"context"
	"errors"
	"provincialig/golimitless/utils"
	"testing"
)

func TestMaxRetrySuccess(t *testing.T) {
	counter := 0

	err := utils.ContinousRetry(context.Background(), 0, 0, func() error {
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

func TestMaxRetryFail(t *testing.T) {
	err := utils.ContinousRetry(context.Background(), 0, 10, func() error {
		return errors.New("")
	})
	if err == nil {
		t.Fail()
	}
}
