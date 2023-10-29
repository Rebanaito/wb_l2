package or

import (
	"sync"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	or := make(chan interface{})
	wg := sync.WaitGroup{}
	wg.Add(len(channels))
	for _, channel := range channels {
		go func(channel <-chan interface{}) {
			defer wg.Done()
			for {
				v, ok := <-channel
				if !ok {
					or <- v
					break
				}
			}
		}(channel)
	}
	go func() {
		wg.Wait()
		close(or)
	}()
	return or
}
