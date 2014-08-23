package main

import (
	"errors"
	"fmt"
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
	result.LengthCA = pointC.Ping(&pointA)
	angA := 0.0
	angB := 0.0
	angC := 0.0

	result.AngleA, angA, _ = GetAngle(result.LengthAB, result.LengthCA, result.LengthBC)
	result.AngleB, angB, _ = GetAngle(result.LengthBC, result.LengthAB, result.LengthCA)
	result.AngleC, angC, _ = GetAngle(result.LengthCA, result.LengthBC, result.LengthAB)
	_, _, _ = angA, angB, angC

	return result
}

var Points []Point

func main() {
	p1 := Point{"P1", false, 1, 1, 0, 0}
	p2 := Point{"P2", false, 2, 4, 0, 0}
	p3 := Point{"P3", false, 3, 3, 0, 0}
	p0 := Point{"P0", true, 4, 2, 0, 0}

	_ = MakeTriangle(p0, p1, p2)
	_ = MakeTriangle(p1, p2, p3)

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
	numerator := (math.Pow(side1, 2)) + (math.Pow(side2, 2) - (math.Pow(side3, 2)))
	denominator := 2 * side1 * side2
	cosine := numerator / denominator
	rads := math.Acos(cosine)
	degrees := 180 / math.Pi * rads
	return rads, degrees, nil
}
