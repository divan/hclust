package hclust

import "testing"

func TestDistance(t *testing.T) {
	a := []float64{5.0, 2.0, 1.0, 4.0, 1.0, 6.0, 2.0}
	b := []float64{1.0, 5.0, 3.0, 5.0, 5.0, 6.0, 1.0}

	tests := []struct {
		dist     DistanceFn
		expected float64
	}{
		{Euclidian, 6.855654600401044},
		{Maximum, 4.0},
		{Canberra, 2.7063492063492065},
		{Manhattan, 15},
		{CosineSimilarity, 0.7862225162829106},
		{Angular, 0.424067912656461},
		{Pearson, 0.8785619869594534},
		{Spearman, 42.94198323981885},
	}

	for _, test := range tests {
		got := test.dist(a, b)
		if got != test.expected {
			t.Helper()
			t.Errorf("expected to get %v, but got %v", test.expected, got)
		}
	}

}
