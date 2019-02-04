package hclust

import (
	"sort"
)

// ranks converts arrays to ranks.
func ranks(a []float64) []float64 {
	sort.Float64s(a)

	// init ranks
	ranks := make([]float64, len(a))
	for i := 0; i < len(ranks); i++ {
		ranks[i] = a[i] + 1
	}

	// fix tied ranks
	tied := make(map[float64]int)
	for _, rank := range ranks {
		tied[rank]++
	}

	for i := 0; i < len(ranks); i++ {
		if tied[ranks[i]] >= 1 {
			// TODO(divan): decrypt this code: sum([...Array(len).keys()].map(el => el + r)) / len
		}
	}
	return ranks
}
