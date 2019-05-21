package hw6test

import "testing"

type testQE struct {
	values []float64
	x1, x2 float64
	ok     bool
}

var testsQE = []testQE{
	{[]float64{1, 1, 1}, 0, 0, false},
	{[]float64{1, 2, 1}, -1, -1, true},
	{[]float64{2, 8, 2}, -0.2679491924311228, -3.732050807568877, true},
}

func TestQuadEquationSet(t *testing.T) {
	for _, data := range testsQE {
		x1, x2, ok := quadEquation(data.values[0], data.values[1], data.values[2])
		if ok != data.ok || x1 != data.x1 || x2 != data.x2 {
			t.Error(
				"For", data.values,
				"expected", data.x1, data.x2, data.ok,
				"got", x1, x2, ok,
			)
		}
	}
}
