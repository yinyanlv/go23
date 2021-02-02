package creational

import (
	"fmt"
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	b := B{A: &A{}}
	b.SetA(1)
	fmt.Println(b.a)
	fmt.Println(b.A.a)
	fmt.Println(b.a)
	if b.a != 2 {
		t.Fatal("abstract factory fatal")
	}
}
