package statistics

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

			for {
				select {
				case <-ctx.Done():
					return
				case msg := <-ch:
					outCh <- msg
				}
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(outCh)
	}()

	return outCh
}
