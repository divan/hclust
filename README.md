# Quick rewrite of Agglomerative hierarchical clustering for Node JS.

[![GoDoc](https://godoc.org/github.com/divan/hclust?status.svg)](https://godoc.org/github.com/divan/hclust)

## Why?

Just to test how quickly I can take [random JS code on internet](https://github.com/Amice13/hclust.js) and rewrite it in Go.

## Installation

```
go get github.com/divan/hclust
```

## Usage

main.go
```
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
```

## Status

This code doesn't work for now.
I haven't implemented a couple of thinks (percent distance, for example), but the most important problem is that JS design is based on the fact that there is no explicit types, so some objects change type on the fly. This design doesn't map well to typed languages, so instead of rewriting it as is in Go, I'd better understand the algrithm better and rewrite it from scratch.

Anyway, it was fun coupld of hours spent.

## Documentation

See [GoDoc](https://godoc.org/github.com/divan/hclust)

## Licence

WTFPL
