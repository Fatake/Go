package main //package name
// package	otroPaquete

import "fmt" // librerías necesarias para importar el paquete

func main() {
	// Variables
	var mensaje string = "Hola mundo"
	easyMessage := "Hola mundo usando :="
	var entero int    //inicia en 0
	var cadena string //inicia en vacio
	var bandera bool  //inicia en fallse
	var complejo complex64 = 10 + 8i
	algo := "variable defida en asignacion"

	//var arreglo []int
	fmt.Printf("Entero %d\n", entero)
	fmt.Printf("cadena %s\n", cadena)
	fmt.Printf("bandera %t\n", bandera)
	fmt.Printf("hello, world %s\n", algo)
	fmt.Println("Complejo ", complejo)
	fmt.Println(mensaje)
	fmt.Println(easyMessage)

	// Puedes comentar usando "//"

	// float64 bits numbers
	a := 10.
	var b float64 = 3
	fmt.Println(a / b)

	//integer numbers
	var c int = 10
	d := 3
	fmt.Println(c / d)

	// boolean
	var x bool = true
	y := false
	fmt.Println(x || y)
	fmt.Println(x && y)
	fmt.Println(!x)

}
