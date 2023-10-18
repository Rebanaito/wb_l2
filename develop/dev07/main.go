package main

import (
	"fmt"
	"sync"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	or := make(chan interface{})
	wg := sync.WaitGroup{}
	wg.Add(len(channels))
	for _, channel := range channels {
		go func(channel <-chan interface{}) {
			for {
				v, ok := <-channel
				if !ok {
					or <- v
				}
			}
			wg.Done()
		}(channel)
	}
	go func() {
		wg.Wait()
		close(or)
	}()
	return or
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
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("Done after %v", time.Since(start))
}
