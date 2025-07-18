package mathutil

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

func TestSub(t *testing.T) {
	result := Sub(2, 3)
	expected := -1
	if result != expected {
		t.Errorf("Sub(2, 3) = %d; want %d", result, expected)
	}
}
