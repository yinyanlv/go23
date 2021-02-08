package structural

import "testing"

func TestAdapter(t *testing.T) {
	adaptee := NewAdaptee()
	adapter := NewAdapter(adaptee)

	if adapter.Request() != "adaptee special request" {
		t.Fatal("adapter error!")
	}
}
