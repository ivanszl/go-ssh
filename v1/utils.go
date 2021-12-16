package v1

import (
	"fmt"
	"time"
)

func WaitForSpecific(f func() (bool, error), maxAttempts int, waitInterval time.Duration) error {
	for i := 0; i < maxAttempts; i++ {
		stop, err := f()
		if err != nil {
			return err
		}
		if stop {
			return nil
		}
		time.Sleep(waitInterval)
	}
	return fmt.Errorf("Maximum number of retries (%d) exceeded", maxAttempts)
}

func WaitFor(f func() (bool, error)) error {
	return WaitForSpecific(f, 30, time.Second*3)
}
