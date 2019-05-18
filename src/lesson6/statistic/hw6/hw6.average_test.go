package hw6

import "testing"

func TestAverage(t *testing.T) {
	v := Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}

func TestSum(t *testing.T) {
	v := Sum([]float64{1, 2})
	if v != 3 {
		t.Error("Expected 3, got ", v)
	}
}
