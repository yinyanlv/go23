// NewXXX函数返回接口时就是简单工厂模式

package creational

import "fmt"

type API interface {
	Say(name string) string
}

func NewAPI(t int) API {
	if t == 1 {
		return &HiAPI{}
	} else if t == 2 {
		return &HelloAPI{}
	}
	return nil
}

type HiAPI struct{}

func (*HiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

type HelloAPI struct{}

func (*HelloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
