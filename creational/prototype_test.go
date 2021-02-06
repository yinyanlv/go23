package creational

import (
	"fmt"
	"testing"
)

var manager *PrototypeManager

type Type1 struct {
	name string
}

func (t *Type1) Clone() Cloneable {
	newT := *t
	return &newT
}

func init() {
	manager = NewPrototypeManager()
	t1 := &Type1{
		name: "type1",
	}
	manager.Set("t1", t1)
}

func TestClone(t *testing.T) {
	t1 := manager.Get("t1")
	fmt.Println(t1)
	t2 := t1.Clone()
	if t1 == t2 {
		t.Fatal("Clone error!")
	}
}

func TestCloneFromManager(t *testing.T) {
	c := manager.Get("t1").Clone()
	t2 := c.(*Type1)
	if t2.name != "type1" {
		t.Fatal("Manager clone error!")
	}
}
