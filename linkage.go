package hclust

import "math"

// LinkageFn defines linkage methods.
type LinkageFn func([]float64) float64

// SingleLinkage implements nearest neighbour method.
func SingleLinkage(a []float64) float64 {
	var min = math.MaxFloat64
	for i := 0; i < len(a); i++ {
		if a[i] < min {
			min = a[i]
		}
	}
	return min
}

// CompleteLinkage implements farthest neighbour method.
func CompleteLinkage(a []float64) float64 {
	var max = -math.MaxFloat64
	for i := 0; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}

// AverageLinkage implements simple average linkage method.
func AverageLinkage(a []float64) float64 {
	return mean(a)
}
