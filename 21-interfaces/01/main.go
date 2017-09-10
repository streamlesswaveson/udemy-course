package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}

func (z Square) area() float64 {
	return z.side * z.side
}

type Circle struct {
	radius float64
}

func (z Circle) area() float64  {
	return math.Pi * (z.radius * z.radius)
}

type Nada struct {

}

type Shape interface {
	area() float64
}

func info(z Shape)  {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main()  {

	s := Square{10}
	c := Circle{10}
	// n := Nada{}
	info(s)
	info(c)
	// throws an error - does not implement Shape
	// info(n)

}
