package math

import "testing"

type testSet struct {
	values []float64
	average float64
}

var testSets = []testSet {
	{ []float64{1,2}, 1.5 },
	{ []float64{1,1,1,1,1,1}, 1 },
	{ []float64{-1,1}, 0 },
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