package ch05

import (
	"testing"
	"time"
)

// 5.3 Heat beat
func DoWork(
	done <-chan interface{},
	nums ...int,
) (<-chan interface{}, <-chan int) {
	heartbeat := make(chan interface{}, 1)
	intStream := make(chan int)
	go func() {
		defer close(heartbeat)
		defer close(intStream)

		time.Sleep(2 * time.Second) // <1>

		for _, n := range nums {
			select {
			case heartbeat <- struct{}{}:
			default:
			}

			select {
			case <-done:
				return
			case intStream <- n:
			}
		}
	}()

	return heartbeat, intStream
}

// P172
func TestDoWork_GeneratesAllNumbersBadCode(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{0, 1, 2, 3, 5}
	_, results := DoWork(done, intSlice...)

	for i, want := range intSlice {
		select {
		case r := <-results:
			if r != want {
				t.Errorf(
					"index %v: expected %v, but received %v,",
					i, want, r,
				)
			}
		case <-time.After(1 * time.Second): // 非決定的な結果になる
			t.Fatal("test timed out")
		}
	}
}
func TestDoWork_GeneratesAllNumbers(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{0, 1, 2, 3, 5}
	heartbeat, results := DoWork(done, intSlice...)

	<-heartbeat // <1>

	i := 0
	for r := range results { // そもそもChannelだからheat beat意味無い気もする。
		if expected := intSlice[i]; r != expected {
			t.Errorf("index %v: expected %v, but received %v,", i, expected, r)
		}
		i++
	}
}
