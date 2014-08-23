package main

import (
	"errors"
	"fmt"
	"log"
)
import "math"

type Point struct {
	Name   string
	IsPrey bool
	X      float64
	Y      float64
	CalcX  float64
	CalcY  float64
}

type Triangle struct {
	LengthAB float64
	LengthBC float64
	LengthCA float64
	AngleA   float64
	AngleB   float64
	AngleC   float64
	PointA   Point
	PointB   Point
	PointC   Point
}

func MakeTriangle(pointA Point, pointB Point, pointC Point) Triangle {
	result := Triangle{PointA: pointA, PointB: pointB, PointC: pointC}
	result.LengthAB = pointA.Ping(&pointB)
	result.LengthBC = pointB.Ping(&pointC)
	result.LengthCA = pointA.Ping(&pointC)
	angA := 0.0
	angB := 0.0
	angC := 0.0
	result.AngleA, angA, _ = GetAngle(result.LengthAB, result.LengthCA, result.LengthBC)
	result.AngleB, angB, _ = GetAngle(result.LengthAB, result.LengthBC, result.LengthCA)
	result.AngleC, angC, _ = GetAngle(result.LengthBC, result.LengthCA, result.LengthAB)
	log.Printf("A:%v + B:%v + C:%v = %v\n", angA, angB, angC, angA+angB+angC)
	return result

	//This doesn't work yet.
	//Figure out why the totals isn't 180!!!
}

var Points []Point

func main() {
	p1 := Point{"P1", false, 1, 1, 0, 0}
	p2 := Point{"P2", false, 2, 4, 0, 0}
	p3 := Point{"P3", false, 5, 2, 0, 0}
	m := make(map[*Point][]float64)
	p0 := Point{"P0", true, 3, 3, 0, 0}

	m[&p0] = append(m[&p0], p0.Ping(&p1))

	log.Printf(p0.ShowDistance(&p1))

	m[&p0] = append(m[&p0], p0.Ping(&p2))
	log.Printf(p0.ShowDistance(&p2))

	m[&p0] = append(m[&p0], p0.Ping(&p3))
	//log.Printf(p0.ShowDistance(&p3))

	//Once we calculate prey, make that the origin and keep it
	//after prey moves away

	m[&p1] = append(m[&p1], p1.Ping(&p0))
	m[&p1] = append(m[&p1], p1.Ping(&p2))
	m[&p1] = append(m[&p1], p1.Ping(&p3))

	log.Printf(p1.ShowDistance(&p2))
	if rads, degs, err := GetAngle(m[&p0][0], m[&p0][1], m[&p0][2]); err != nil {
		log.Printf("Error %v", err)
	} else {
		log.Printf("Hyp Angle r:%v d:%v", rads, degs)
	}

	m[&p2] = append(m[&p2], p2.Ping(&p0))
	m[&p2] = append(m[&p2], p2.Ping(&p1))
	m[&p2] = append(m[&p2], p2.Ping(&p3))

	m[&p3] = append(m[&p3], p3.Ping(&p0))
	m[&p3] = append(m[&p3], p3.Ping(&p1))
	m[&p3] = append(m[&p3], p3.Ping(&p2))

	_ = MakeTriangle(p0, p1, p2)

}

func (p *Point) Ping(otherPoint *Point) float64 {
	distance := math.Sqrt(math.Pow(p.X-otherPoint.X, 2) + math.Pow(p.Y-otherPoint.Y, 2))
	return distance
}

func (p *Point) ShowDistance(otherPoint *Point) string {
	distance := math.Sqrt(math.Pow(p.X-otherPoint.X, 2) + math.Pow(p.Y-otherPoint.Y, 2))
	return fmt.Sprintf("Distance between %v and %v is %v.", p.Name, otherPoint.Name, distance)
}

func GetAngle(side1 float64, side2 float64, side3 float64) (float64, float64, error) {
	if side1 <= 0 || side2 <= 0 || side3 <= 0 {
		return 0, 0, errors.New("A, B, & C must all be greater than 0")
	}
	numerator := (math.Pow(side3, 2)) - (math.Pow(side1, 2) - (math.Pow(side2, 2)))
	denominator := -2 * side1 * side2
	cosine := numerator / denominator
	rads := math.Acos(cosine)
	degrees := 180 / math.Pi * rads
	return rads, degrees, nil
}
