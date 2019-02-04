package hclust

func distanceMatrix(arrays [][]float64, dist DistanceFn, NAPairwise bool) []Cluster {
	var ret []Cluster
	for i := 0; i < len(arrays); i++ {
		for j := i + 1; j < len(arrays); j++ {
			a := arrays[i]
			b := arrays[j]
			if len(a) == 0 || len(b) == 0 {
				panic("invalid input data")
			}

			arrayDistance := dist(a, b)
			cluster := NewCluster(a, b, arrayDistance)
			ret = append(ret, cluster)
		}
	}

	return ret
}
