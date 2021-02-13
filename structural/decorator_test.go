package structural

import "testing"

func TestDecorator(t *testing.T) {
	var c CalcComponent = &ConcreteComponent{}
	c = WrapAddDecorator(c, 10)
	c = WrapMulDecorator(c, 8)

	res := c.Calc()

	if res != 80 {
		t.Fail()
	}
}
