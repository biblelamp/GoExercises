package hw6test

import "testing"

type testpair struct {
	values  []float64
	average float64
	total   float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5, 3},
	{[]float64{1, 1, 1, 1, 1, 1}, 1, 6},
	{[]float64{-1, 1}, 0, 0},
}

func TestAverageSet(t *testing.T) {
	for _, pair := range tests {
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

func TestSumSet(t *testing.T) {
	for _, pair := range tests {
		v := Sum(pair.values)
		if v != pair.total {
			t.Error(
				"For", pair.values,
				"expected", pair.total,
				"got", v,
			)
		}
	}
}
