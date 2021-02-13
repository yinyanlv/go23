// 装饰模式使用对象组合的方式动态改变或增加对象行为

package structural

type CalcComponent interface {
	Calc() int
}

type ConcreteComponent struct {
}

func (*ConcreteComponent) Calc() int {
	return 0
}

type MulDecorator struct {
	CalcComponent
	num int
}

func WrapMulDecorator(c CalcComponent, num int) CalcComponent {
	return &MulDecorator{
		CalcComponent: c,
		num:           num,
	}
}

func (d *MulDecorator) Calc() int {
	return d.CalcComponent.Calc() * d.num
}

type AddDecorator struct {
	CalcComponent
	num int
}

func WrapAddDecorator(c CalcComponent, num int) CalcComponent {
	return &AddDecorator{
		CalcComponent: c,
		num:           num,
	}
}

func (d *AddDecorator) Calc() int {
	return d.CalcComponent.Calc() + d.num
}
