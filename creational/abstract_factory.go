package creational

type A struct {
	a int
}

func (a *A) SetA(num int) {
	a.a = num
}

type B struct {
	*A
}
