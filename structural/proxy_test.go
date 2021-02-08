package structural

import "testing"

func TestProxy(t *testing.T) {
	sub := Proxy{}
	res := sub.Do()
	if res != "pre:real:after" {
		t.Fatal("proxy error!")
	}
}
