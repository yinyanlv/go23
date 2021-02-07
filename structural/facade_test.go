package structural

import "testing"

func TestFacade(t *testing.T) {
	api := NewAPI()
	str := api.Test()
	if str != "a module ap1\nb module api" {
		t.Fatal("facade error")
	}
}
