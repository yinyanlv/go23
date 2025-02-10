package stability

import (
	"fmt"
	"time"
)

// 重试模式
func Retry(attempts int, duration time.Duration, fn func() error) error {
	for i := 0; ; i++ {
		err := fn()
		if err == nil {
			return nil
		}

		if i > attempts {
			return fmt.Errorf("超过最大重试次数，%v", err)
		}
		time.Sleep(duration)
		// 指数退避策略
		duration *= 2
	}
}
