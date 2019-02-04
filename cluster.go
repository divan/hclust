package hclust

// Cluster represents the single cluster
// information, including its data and distance.
type Cluster struct {
	Elements [][]float64 `json:"elements"`
	Distance float64     `json:"distance"`
	Indices  []int       `json:"indices,omitempty"`
}

// NewCluster creates a new Cluster.
func NewCluster(a, b []float64, dist float64) Cluster {
	return Cluster{
		Elements: [][]float64{a, b},
		Distance: dist,
	}
}
