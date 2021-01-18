package main //package name
// package	otroPaquete

import "fmt" // librerías necesarias para importar el paquete

func main() {
	// Variables
	var mensaje string = "Hola mundo"
	easyMessage := "Hola mundo usando :="
	fmt.Println(mensaje)
	fmt.Println(easyMessage)

	// Puedes comentar usando "//"

	// float numbers
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

	//Lógica privada
	privada()

	//Lógica Publica
	Publica()

	//Función "defer"
	printHelloWorld()
}

// func types
func privada() {
	fmt.Println("Ejecutar lógica que no necesita ser exportada (pertenece solo a este paquete)")
}

/**
* Compentario
 */
func Publica() {
	/**
	* Comentario Publico
	 */
	fmt.Println("Lógica que quiero exportar a otros paquetes")
}

// defer
func printHelloWorld() {
	defer fmt.Println("World!")
	fmt.Println("Hello")

	fmt.Println("defer ejecuta un fragmento de código hasta que la función haya terminado")
}
