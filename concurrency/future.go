package concurrency

// 未来模式
type Future struct {
	result interface{}
	err    error
	done   chan struct{}
}

func Async(fn func() (interface{}, error)) *Future {
	f := &Future{
		done: make(chan struct{}),
	}

	go func() {
		f.result, f.err = fn()
		close(f.done)
	}()
	return f
}

func (f *Future) Get() (interface{}, error) {
	<-f.done
	return f.result, f.err
}
