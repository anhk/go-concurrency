package go_concurrency

import (
	"context"
	"log"
	"sync"
	"time"
)

type CallbackFunc func(args ...interface{}) error

func Run(cb CallbackFunc, timeout time.Duration, args ...interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	done := make(chan error, 1)
	go func() {
		done <- cb(args)
	}()

	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func RunBatch(cb CallbackFunc, timeout time.Duration, count int, args ...interface{}) error {
	wg := sync.WaitGroup{}
	wg.Add(count)

	for i := 0; i < count; i++ {
		go func() {
			defer func() {
				if p := recover(); p != nil {
					log.Println("oops, panic:", p)
				}
			}()
			defer wg.Done()
			Run(cb, timeout, args)
		}()
	}

	wg.Wait()
	return nil
}
