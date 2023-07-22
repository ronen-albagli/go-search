package timer

import (
	"fmt"
	"time"
)

type Timer struct {
	since time.Time
}

func (t *Timer) Start() {
	t.since = time.Now()
}

func (t *Timer) Finish() {

	duration := time.Since(t.since).Milliseconds()

	fmt.Printf("Execution time: %d milliseconds\n", duration)
}
