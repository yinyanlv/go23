package concurrency

// 扇出模式
func FanOut(input <-chan int, numWorkers int, processor func(int)) {
	for i := 0; i < numWorkers; i++ {
		go func() {
			for v := range input {
				processor(v)
			}
		}()
	}
}
