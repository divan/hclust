package hclust

import (
	"math"
)

// DistanceFn describes distance function.
type DistanceFn func(a, b []float64) float64

// Euclidian implements Euclidian distance formula for two points
// in an Euclidian space. a and b should be of the same length.
//
// See https://en.wikipedia.org/wiki/Euclidean_distance
var Euclidian DistanceFn = func(a, b []float64) float64 {
	var sum float64
	for i := 0; i < len(a); i++ {
		sum += math.Pow(a[i]-b[i], 2)
	}

	return math.Sqrt(sum)
}

// Maximum implements Maximum distance formula.
// a and b should be of the same length.
var Maximum DistanceFn = func(a, b []float64) float64 {
	var max float64
	for i := 0; i < len(a); i++ {
		dist := math.Abs(a[i] - b[i])
		if dist > max {
			max = dist
		}
	}

	return max
}

// Canberra implements Canberra distance formula.
// a and b should be of the same length.
//
// See https://en.wikipedia.org/wiki/Canberra_distance
var Canberra DistanceFn = func(a, b []float64) float64 {
	var sum float64
	for i := 0; i < len(a); i++ {
		sum += math.Abs(a[i]-b[i]) / (math.Abs(a[i]) + math.Abs(b[i]))
	}

	return sum
}

// Manhattan implements Manhattan distance formula.
// a and b should be of the same length.
//
// See https://en.wikipedia.org/wiki/Taxicab_geometry
var Manhattan DistanceFn = func(a, b []float64) float64 {
	var sum float64
	for i := 0; i < len(a); i++ {
		sum += math.Abs(a[i] - b[i])
	}

	return sum
}

// Percent implements percent distance formula.
// a and b should be of the same length.
var Percent DistanceFn = func(a, b []float64) float64 {
	var sum float64
	for i := 0; i < len(a); i++ {
		sum += math.Abs(a[i] - b[i])
	}

	return sum
}

// CosineSimilarity implements cosine similarity formula.
// a and b should be of the same length.
//
// See https://en.wikipedia.org/wiki/Cosine_similarity
var CosineSimilarity DistanceFn = func(a, b []float64) float64 {
	num := dotProduct(a, b)
	denom := math.Sqrt(dotProduct(a, a) * dotProduct(b, b))
	if denom == 0 {
		panic("denominator can't be zero")
	}
	return num / denom
}

// Angular implements angular distance formula.
// a and b should be of the same length.
var Angular DistanceFn = func(a, b []float64) float64 {
	cs := CosineSimilarity(a, b)
	return 2 * math.Acos(cs) / math.Pi
}

// Pearson implements Pearson's distance formula.
// a and b should be of the same length.
//
// See https://en.wikipedia.org/wiki/Pearson_correlation_coefficient#Pearson%27s_distance
var Pearson DistanceFn = func(a, b []float64) float64 {
	num := dotProduct(residuals(a), residuals(b))
	denom := sd(a) * sd(b) * float64(len(a))
	if denom == 0 {
		panic("denominator can't be zero")
	}
	pearsonCoeff := num / denom
	return 1 - pearsonCoeff
}

// Spearman implements Spearman's distance formula.
// a and b should be of the same length.
var Spearman DistanceFn = func(a, b []float64) float64 {
	ra := ranks(a)
	rb := ranks(b)
	dist := Euclidian(ra, rb)
	length := float64(len(a))
	return 1 - 6*(dist/math.Pow(length, 3)-length) // ??
}
