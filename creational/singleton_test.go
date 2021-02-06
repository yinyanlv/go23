package creational

import (
	"sync"
	"testing"
)

func TestSingleton(t *testing.T) {
	ins1 := GetInstance()
	ins2 := GetInstance()

	if ins1 != ins2 {
		t.Fatal("Singleton error!")
	}
}

func TestParallelSingleton(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(count)
	instances := [count]*Singleton{}
	for i := 0; i < count; i++ {
		go func(index int) {
			instances[index] = GetInstance()
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 1; i < count; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("Parallel singleton error!")
		}
	}
}
