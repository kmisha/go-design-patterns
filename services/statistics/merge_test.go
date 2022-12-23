package statistics

import (
	"context"
	"testing"
	"time"
)

func generate(data string, amount int) <-chan string {
	channel := make(chan string)

	go func() {
		for i := 0; i < amount; i++ {
			channel <- data
			time.Sleep(time.Duration(100 * time.Millisecond))

		}

		close(channel)
	}()
	return channel
}
func TestMergeData(t *testing.T) {
	t.Run("we can push data from channels in to result channel", func(t *testing.T) {
		ch1 := generate("first", 5)
		ch2 := generate("second", 5)
		ctx, cansel := context.WithTimeout(context.TODO(), time.Second)
		defer cansel()

		want := map[string]int{
			"first":  0,
			"second": 0,
		}

		got := MergeData(ctx, ch1, ch2)

		for msg := range got {
			want[msg] = want[msg] + 1
		}

		if want["first"] != 5 || want["second"] != 5 {
			t.Fatalf("wont 5 first and 5 second; got %d first and %d second", want["first"], want["second"])
		}
	})

	t.Run("we should get no more message from result channel than sent", func(t *testing.T) {
		ch1 := generate("first", 5)
		ch2 := generate("second", 10)
		ctx, cansel := context.WithTimeout(context.TODO(), time.Second)
		defer cansel()
		got := 0
		want := 15

		for m := range MergeData(ctx, ch1, ch2) {
			if len(m) > 1 {
				got += 1
			}
		}

		if got != want {
			t.Fatalf("want %d messages but got %d", want, got)
		}

	})

	t.Run("we must respect context", func(t *testing.T) {
		t.Skip("TODO rewrite this test")
		ch1 := generate("first", 5)
		ch2 := generate("second", 10)
		ctx, _ := context.WithTimeout(context.TODO(), time.Millisecond)

		ch := MergeData(ctx, ch1, ch2)
		time.Sleep(2 * time.Millisecond)
		_, ok := (<-ch)

		if ok {
			t.Fatalf("channel isn't closed when ctx done")
		}
	})
}
