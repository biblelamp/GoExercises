package math

import (
	"math"
	"testing"
)

type testSet struct {
	values []float64
	average float64
	min float64
}

var testSets = []testSet {
	{ []float64{1,2}, 1.5, 1 },
	{ []float64{1,1,1,1,1,1}, 1, 1 },
	{ []float64{-1,1}, 0, -1 },
	{ []float64{}, 0, math.NaN() },
}

func TestAverageOne(t *testing.T) {
	var v float64
	v = Average([]float64{1,2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}

func TestAverage(t *testing.T) {
	for _, pair := range testSets {
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}

func TestMin(t *testing.T) {
	for _, pair := range testSets {
		v := Min(pair.values)
		if len(pair.values) == 0 && math.IsNaN(v) {
			continue
		}
		if v != pair.min {
			t.Error(
				"For", pair.values,
				"expected", pair.min,
				"got", v,
			)
		}
	}
}