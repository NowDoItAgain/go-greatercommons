package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	a := s.length * s.length
	return a
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {

	earth := circle{
		radius: 45.4,
	}

	box := square{
		length: 12.2,
	}

	info(earth)
	info(box)
}
