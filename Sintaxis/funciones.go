package main //package name
// package	otroPaquete

import "fmt" // librerías necesarias para importar el paquete

func main() {
	fmt.Println("Tipos de funciones en Go, Publicas, privada y defer")

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
