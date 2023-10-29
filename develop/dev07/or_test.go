package or

import (
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	done := false
	go func() {
		<-or(
			sig(2*time.Hour),
			sig(5*time.Minute),
			sig(1*time.Second),
			sig(1*time.Hour),
			sig(1*time.Minute),
		)
		done = true
	}()
	time.Sleep(2 * time.Second)
	if !done {
		t.Fatal(`Done channel did not terminate`)
	}
}
