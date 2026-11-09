package utils

import (
	"context"
	"time"
)

func Interval(ctx context.Context, fn func(), ms time.Duration) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(ms):
				fn()
			}
		}
	}()
}
