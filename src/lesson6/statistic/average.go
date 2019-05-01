package statistic

// How to create package:
// save to %GOPATH%/src/statistics
// move to this folder and make:
// > go install

func Average(xs []float64) float64 {
    total := float64(0)
    for _, x := range xs {
        total += x
    }
    return total / float64(len(xs))
}