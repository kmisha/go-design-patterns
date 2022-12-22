package merge

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
		want := map[string]int{
			"first":  0,
			"second": 0,
		}

		got := MergeData(context.TODO(), ch1, ch2)

		for msg := range got {
			want[msg] = want[msg] + 1
		}

		if want["first"] != 5 || want["second"] != 5 {
			t.Fatalf("wont 5 first and 5 second; got %d first and %d second", want["first"], want["second"])
		}

	})
}
