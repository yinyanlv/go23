package creational

import "testing"

func TestType1(t *testing.T) {
	api := NewAPI(1)
	s := api.Say("Bugong")
	if s != "Hi, Bugong" {
		t.Fatal("Type1 test failed!")
	}
}

func TestType2(t *testing.T) {
	api := NewAPI(2)
	s := api.Say("Bugong")
	if s != "Hello, Bugong" {
		t.Fatal("Type2 test failed!")
	}
}
