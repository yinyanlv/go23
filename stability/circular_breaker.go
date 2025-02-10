package stability

import (
	"errors"
	"sync"
	"time"
)

// 断路器模式
type CircularBreaker struct {
	failures     int
	maxFailures  int
	resetTimeout time.Duration
	state        int // 0: 关闭, 1: 打开, 2: 半开
	mutex        sync.Mutex
	lastFailure  time.Time
}

func NewCircularBreaker(maxFailures int, timeout time.Duration) *CircularBreaker {
	return &CircularBreaker{
		maxFailures:  maxFailures,
		resetTimeout: timeout,
	}
}

func (cb *CircularBreaker) Execute(action func() error) error {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()
	// 断路器已开，上次失败的时间已超出重置断路器时间
	if cb.state == 1 && time.Since(cb.lastFailure) > cb.resetTimeout {
		// 设置为半开状态
		cb.state = 2
	}
	if cb.state == 1 {
		return errors.New("circular breaker open")
	}
	err := action()
	if err != nil {
		cb.failures++
		cb.lastFailure = time.Now()
		// 打开断路器
		if cb.failures >= cb.maxFailures {
			cb.state = 1
		}
		return err
	}
	// 半开情况下重置断路器
	if cb.state == 2 {
		cb.reset()
	}
	return nil
}

func (cb *CircularBreaker) reset() {
	cb.failures = 0
	cb.state = 0
}
