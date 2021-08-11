package terminal

import "testing"

func TestWidth(t *testing.T) {
	result := Width()
	if result < 0 {
		t.Fatalf("Negative Width: Width() = %v", result)
	} else if result > 1000 {
		t.Fatalf("Width is bigger than expected: Width() = %v", result)
	}
}