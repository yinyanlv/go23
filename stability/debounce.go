package stability

import (
	"sync"
	"time"
)

// 防抖模式
func Debounce(d time.Duration, f func()) func() {
	var timer *time.Timer
	var mu sync.Mutex

	return func() {
		mu.Lock()
		defer mu.Unlock()

		if timer != nil {
			timer.Stop()
		}

		timer = time.AfterFunc(d, f)
	}
}
