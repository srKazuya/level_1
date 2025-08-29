package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func Distance(one, two Point) float64 {
	answ := math.Sqrt((math.Pow(two.x - one.x, 2)) + (math.Pow(two.y-one.y, 2)))
	return answ
}

func main() {
	one := NewPoint(2, 4)
	two := NewPoint(3, 5)
	
	answ := Distance(one, two)
	fmt.Println(answ)
}
