package stability

import "time"

// 节流模式
type Throttler struct {
	tokens chan struct{}
}

func NewThrottler(rate int) *Throttler {
	t := &Throttler{
		tokens: make(chan struct{}, rate),
	}
	for i := 0; i < rate; i++ {
		t.tokens <- struct{}{}
	}

	go t.refill(time.Second)
	return t
}

// 定时检测并放入令牌
func (t *Throttler) refill(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		for i := 0; i < cap(t.tokens); i++ {
			select {
			case t.tokens <- struct{}{}:
			default:
			}
		}
	}
}

func (t *Throttler) Allow() bool {
	select {
	case <-t.tokens:
		return true
	default:
		return false
	}
}
