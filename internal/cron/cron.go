package cron

import (
	"sync"
	"time"
)

func RunAndRepeat(duration time.Duration, f func()) {
	var lock sync.Mutex
	timer := time.NewTicker(duration)
	defer timer.Stop()

	job := func() {
		lock.Lock()
		defer lock.Unlock()
		f()
	}

    // Run job directly
    go job()

    // Repeat once timer ticks
	for {
		select {
		case <-timer.C:
			go job()
		}
	}
}
