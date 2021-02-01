package creational

import "testing"

func compute(f OperatorFactory, a, b int) int {
	op := f.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

func TestOperator(t *testing.T) {
	var f OperatorFactory
	f = PlusOperatorFactory{}
	if compute(f, 1, 2) != 3 {
		t.Fatal("Error: plus operator factory")
	}

	f = MinusOperatorFactory{}
	if compute(f, 3, 1) != 2 {
		t.Fatal("Error: minus operator factory")
	}

}
