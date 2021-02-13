// API为facade模块的外观接口，大部分代码使用此接口简化对facade类的访问

package structural

import "fmt"

type API interface {
	Test() string
}

func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func (a *apiImpl) Test() string {
	strA := a.a.TestA()
	strB := a.b.TestB()

	return fmt.Sprintf("%s\n%s", strA, strB)
}

func NewAModuleAPI() AModuleAPI {
	return &aModuleAPI{}
}

type AModuleAPI interface {
	TestA() string
}

type aModuleAPI struct {
}

func (a *aModuleAPI) TestA() string {
	return "a module api"
}

func NewBModuleAPI() BModuleAPI {
	return &bModuleAPI{}
}

type BModuleAPI interface {
	TestB() string
}

type bModuleAPI struct{}

func (b *bModuleAPI) TestB() string {
	return "b module api"
}
