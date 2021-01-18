package main

import (
	"fmt"
	"math"
)

type geometry interface {
	calculateArea() float64
	calculatePerimeter() float64
}

func printArea(g geometry) {
	fmt.Println(g.calculateArea())
}
func printPerimeter(g geometry) {
	fmt.Println(g.calculatePerimeter())
}

// Circulo
type circle struct {
	radio float64
}

func (circle *circle) calculateArea() float64 {
	return 3.14159 * math.Pow(circle.radio, 2)
}
func (circle *circle) calculatePerimeter() float64 {
	return 3.14159 * (circle.radio * 2)
}

// Cuadrado

type square struct {
	side float64
}

func (square *square) calculateArea() float64 {
	return math.Pow(square.side, 2)
}
func (square *square) calculatePerimeter() float64 {
	return square.side * 4
}

func main() {
	var square = &square{side: 3}
	fmt.Println("----- Cuadrado implementando Geometrico -----")
	fmt.Println("Largo: ", square.side)
	fmt.Println("Area: ", geometry.calculateArea(square))
	fmt.Println("Perimetro: ", geometry.calculatePerimeter(square))

	var circle = &circle{radio: 3}
	fmt.Println("----- Circulo implementando Geometrico -----")
	fmt.Println("Radio: ", circle.radio)
	fmt.Println("Area: ", geometry.calculateArea(circle))
	fmt.Println("Perimetro: ", geometry.calculatePerimeter(circle))
}
