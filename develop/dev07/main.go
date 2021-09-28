package main

import (
	"context"
	"fmt"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {

	ctx, cancel := context.WithCancel(context.Background())

	res := make(chan interface{})

	for _, ch := range channels {

		go func(ctx context.Context, c <-chan interface{}) {
			select {
			case <-c:
				cancel()
				res <- struct{}{}
			case <-ctx.Done():
				res <- struct{}{}
			}
		}(ctx, ch)
	}

	return res
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(2*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("fone, after %v\n", time.Since(start))
}
