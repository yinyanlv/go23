// 代理模式用于延迟处理操作或者在进行实际操作前后进行其它处理

package structural

type Subject interface {
	Do() string
}

type RealSubject struct {
}

func (RealSubject) Do() string {
	return "real"
}

type Proxy struct {
	real RealSubject
}

func (p Proxy) Do() string {
	res := "pre:"
	res += p.real.Do()
	res += ":after"
	return res
}
