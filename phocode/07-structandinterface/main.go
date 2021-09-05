package main

import (
	"fmt"
	"math"
)

func main() {
	c := Circle{0, 0, 5}
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(circleArea(c))
	fmt.Println(rectangleArea(r))
	fmt.Println(circleArea2(&c))
	fmt.Println(c.area())
	fmt.Println(r.area())

}

type Shape interface {
	area() float64
}

type Circle struct {
	x float64
	y float64
	r float64
}

type Rectangle struct {
	x1 float64
	y1 float64
	x2 float64
	y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func rectangleArea(r Rectangle) float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func circleArea2(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}
