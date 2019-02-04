package hclust

import "math"

// LinkageFn defines linkage methods.
type LinkageFn func([]float64) float64

// SingleLinkage implements nearest neighbour method.
var SingleLinkage = func(a []float64) float64 {
	var min = math.MaxFloat64
	for i := 0; i < len(a); i++ {
		if a[i] < min {
			min = a[i]
		}
	}
	return min
}

// CompleteLinkage implements farthest neighbour method.
var CompleteLinkage = func(a []float64) float64 {
	var max = -math.MaxFloat64
	for i := 0; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}

// AverageLinkage implements simple average linkage method.
var AverageLinkage = func(a []float64) float64 {
	return mean(a)
}
