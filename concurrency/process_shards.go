package concurrency

import "sync"

// 分片模式
func ProcessShards(data []int, shardSize int, process func([]int) []int) []int {
	var wg sync.WaitGroup
	results := make(chan []int)

	for i := 0; i < len(data); i += shardSize {
		end := i + shardSize
		if end > len(data) {
			end = len(data)
		}
		shard := data[i:end]
		wg.Add(1)
		go func(s []int) {
			defer wg.Done()
			results <- process(s)
		}(shard)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var final []int
	for res := range results {
		final = append(final, res...)
	}
	return final
}
