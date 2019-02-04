package hclust

import "math"

// dotProduct naively calculates dot product for two vectors.
func dotProduct(a, b []float64) float64 {
	var sum float64
	for i := 0; i < len(a); i++ {
		sum += a[i] * b[i]
	}

	return sum
}

// residuals calculates residuals for the array.
func residuals(a []float64) []float64 {
	ret := make([]float64, len(a))
	m := mean(a)
	for i := 0; i < len(a); i++ {
		ret[i] = a[i] - m
	}
	return ret
}

// mean calculates mean value for the array.
func mean(a []float64) float64 {
	var sum float64
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum / float64(len(a))
}

// variance returns variance for the array.
func variance(a []float64) float64 {
	var sum float64

	m := mean(a)
	for i := 0; i < len(a); i++ {
		sum += math.Pow(a[i]-m, 2)
	}

	return sum / float64(len(a))
}

// sd returns standart deviation for the array.
func sd(a []float64) float64 {
	return math.Sqrt(variance(a))
}
