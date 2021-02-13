// 工厂方法模式使用子类的方式延迟生成对象到子类中实现

package creational

type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

type OperatorFactory interface {
	Create() Operator
}

type BaseOperator struct {
	a, b int
}

func (o *BaseOperator) SetA(num int) {
	o.a = num
}

func (o *BaseOperator) SetB(num int) {
	o.b = num
}

type PlusOperatorFactory struct{}
type PlusOperator struct {
	*BaseOperator
}

func (o PlusOperator) Result() int {
	return o.a + o.b
}

func (o PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		BaseOperator: &BaseOperator{},
	}
}

type MinusOperatorFactory struct{}
type MinusOperator struct {
	*BaseOperator
}

func (o MinusOperator) Result() int {
	return o.a - o.b
}

func (f MinusOperatorFactory) Create() Operator {
	return MinusOperator{
		BaseOperator: &BaseOperator{},
	}
}
