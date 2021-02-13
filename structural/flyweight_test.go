package structural

import "testing"

func TestFlyweight(t *testing.T) {
	viewer1 := NewImageViewer("image1.png")
	viewer2 := NewImageViewer("image1.png")
	if viewer1.ImageFlyweight != viewer2.ImageFlyweight {
		t.Fail()
	}
}
