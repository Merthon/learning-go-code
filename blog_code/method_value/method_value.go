package main

import (
	"fmt"
	"math"
)

type Point struct{X, Y float64}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main()  {
	p := Point{1, 2}
	q := Point{4, 6}

	distanceFormP := p.Distance
	fmt.Println(distanceFormP(q))
	fmt.Println(p.Distance(q))

	distanceFormQ := q.Distance
	fmt.Println(distanceFormQ(p))
	fmt.Println(p.Distance(p))
}