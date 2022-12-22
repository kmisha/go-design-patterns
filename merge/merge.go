package merge

import (
	"context"
	"sync"
)

func MergeData[T interface{}](ctx context.Context, channels ...<-chan T) <-chan T {
	outCh := make(chan T)

	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, c := range channels {
		go func(ch <-chan T) {
			defer wg.Done()

			for m := range ch {
				outCh <- m
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(outCh)
	}()

	return outCh
}
