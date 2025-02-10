package concurrency

import "sync"

// 扇入模式
func FanIn(inputs ...chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range inputs {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
