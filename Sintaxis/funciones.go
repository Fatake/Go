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

	r1, r2 := dobleRetorno(2, 1)
	fmt.Println("Numeros Retornados", r1, r2)
}

func unParametro(numero int) int {
	fmt.Println("Numero dado: ", numero)
	return numero + 1
}

func dobleRetorno(a, b int) (c, d int) {
	fmt.Println("Numeros dados", a, b)
	return a + b, b - a
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
