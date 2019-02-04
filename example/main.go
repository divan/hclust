package main

import (
	"fmt"

	"github.com/divan/hclust"
)

func main() {
	var arrays = [][]float64{
		{5.0, 2.0, 1.0, 4.0, 1.0, 6.0, 2},
		{1.0, 5.0, 3.0, 5.0, 5.0, 6.0, 1},
		{6.0, 2.0, 7.0, 7.0, 5.0, 6.0, 7},
		{1.0, 2.0, 3.0, 4.0, 6.0, 6.0, 7},
	}

	clusters := hclust.Build(arrays, hclust.Euclidian, hclust.AverageLinkage, false)

	fmt.Println(clusters)
}
