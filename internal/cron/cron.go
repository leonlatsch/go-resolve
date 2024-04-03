package cron

import (
	"sync"
	"time"
)

func Repeat(duration time.Duration, f func()) {
	var lock sync.Mutex
	timer := time.NewTicker(duration)
	defer timer.Stop()

	job := func() {
		lock.Lock()
		defer lock.Unlock()
		f()
	}

	// Repeat once timer ticks
	for {
		select {
		case <-timer.C:
			go job()
		}
	}
}
