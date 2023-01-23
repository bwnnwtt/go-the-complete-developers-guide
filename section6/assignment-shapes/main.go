package main

import "fmt"

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

type shape interface {
	getArea() float64
}

// test that the square, triangle and shape implemented works
func main() {

	// initialize a square and triangle
	sq := square{5.0}
	tri := triangle{height: 10, base: 20}

	// output the area of square and triangle
	printArea(sq)
	printArea(tri)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
