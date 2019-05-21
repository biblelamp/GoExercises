package hw6test

func Average(xs []float64) float64 {
	return Sum(xs) / float64(len(xs))
}

func Sum(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total
}
