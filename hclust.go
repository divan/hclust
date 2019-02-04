package hclust

import (
	"math"
)

// Build distance matrix for arrays, using given distance function, linkage method and NA strategy.
func Build(arrays [][]float64, dist DistanceFn, linkage LinkageFn, NAPairwise bool) []Cluster {
	distanceMatrix := distanceMatrix(arrays, dist, NAPairwise)

	elements := make([][]float64, len(arrays))
	for i := range arrays {
		elements[i] = []float64{float64(i)}
	}

	var ret []Cluster
	for len(elements) != 1 {
		var clusters []Cluster
		for i := 0; i < len(elements); i++ {
			for j := i + 1; j < len(elements); j++ {
				a := flatten(elements[i])
				b := flatten(elements[j])

				var distances []float64
				for _, cluster := range distanceMatrix {
					if common(cluster.Elements[0], a) > 0 && common(cluster.Elements[0], b) > 0 {
						distances = append(distances, cluster.Distance)
					}

				}

				distance := linkage(distances)

				cluster := NewCluster(a, b, distance)
				cluster.Indices = []int{i, j}
				clusters = append(clusters, cluster)
			}
		}

		var (
			minDistance    = math.MaxFloat64
			closestCluster Cluster
		)
		for _, cluster := range clusters {
			if cluster.Distance < minDistance {
				minDistance = cluster.Distance
				closestCluster = cluster
			}
		}

		elements = filterClustersWithIndices(elements, closestCluster.Indices)
		ret = append(ret, closestCluster)
		elements = append(elements, closestCluster.Elements[0])
	}

	return ret
}

func flatten(a []float64) []float64 {
	return a // TODO: implement this
}

// common returns number of common items in two arrays.
//
// In JS code it was 'intersection' and returned the whole array,
// but as caller code just checks for len, it makes sense to
// just return len here.
func common(a, b []float64) int {
	var ret []float64
	for _, valueA := range a {
		for _, valueB := range b {
			if valueB == valueA {
				ret = append(ret, valueA)
				continue
			}
		}
	}
	return len(ret)
}

// attempt to make this JS code:
// `elements = elements.filter((value, index) => !cluster.indicies.includes(index))`
// readable.
func filterClustersWithIndices(elements [][]float64, indices []int) [][]float64 {
	var ret [][]float64
	for index, value := range elements {
		var includes bool
		for _, idx := range indices {
			if index == idx {
				includes = true
			}
		}

		if !includes {
			ret = append(ret, value)
		}
	}
	return ret
}
