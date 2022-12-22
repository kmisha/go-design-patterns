package merge

import (
	"context"
	"sync"
)

func MergeData(ctx context.Context, channels ...<-chan string) <-chan string {
	outCh := make(chan string)

	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, c := range channels {
		go func(ch <-chan string) {
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
