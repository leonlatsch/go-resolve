package cron

import (
	"sync"
	"time"
)

func Repeat(duration time.Duration, f func()) {
	var lock sync.Mutex

	job := func() {
		lock.Lock()
		defer lock.Unlock()
		f()
	}

	for {
		time.Sleep(duration)
		job()
	}
}
