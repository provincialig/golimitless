package utils_test

import (
	"context"
	"errors"
	"provincialig/golimitless/utils"
	"testing"
)

func TestUtils(t *testing.T) {
	t.Run("MaxRetrySuccess", func(t *testing.T) {
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
	})

	t.Run("MaxRetryFail", func(t *testing.T) {
		err := utils.ContinousRetry(context.Background(), 0, 10, func() error {
			return errors.New("")
		})
		if err == nil {
			t.Fail()
		}
	})
}
