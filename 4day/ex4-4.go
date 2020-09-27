package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (c circle) area() {
	fmt.Println(math.Pi * c.radius * c.radius)
}

func (s square) area() {
	fmt.Println(s.length * s.length)
}

type shape interface {
	area()
}

func info(s shape) {
	s.area()
}

func main() {
	c := circle{
		5,
	}

	s := square{
		4,
	}

	c.area()
	s.area()

	info(c)
	info(s)
}
