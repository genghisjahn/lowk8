package main

import "log"
import "math"

type Point struct {
	IsPrey bool
	X      float64
	Y      float64
	CalcX  float64
	CalcY  float64
}

var Points []Point

func main() {
	p1 := Point{false, 5, 5, 0, 0}
	p2 := Point{false, 6, 6, 0, 0}
	p3 := Point{false, 7, 7, 0, 0}

	p0 := Point{true, 9, 9, 0, 0}
	dbtween := p0.Ping(&p1)
	log.Printf("Distance between P0 to P1: %v\n", dbtween)
	dbtween = p0.Ping(&p2)
	log.Printf("Distance between P0 to P2: %v\n", dbtween)
	dbtween = p0.Ping(&p3)
	log.Printf("Distance between P0 to P3: %v\n", dbtween)

	//Let's treat p0 as the origin every time we recalulate
}

func (p *Point) Ping(otherPoint *Point) float64 {
	distance := math.Sqrt(math.Pow(p.X-otherPoint.X, 2) + math.Pow(p.Y-otherPoint.Y, 2))
	return distance
}
