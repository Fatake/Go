package main

import "fmt"

func main() {
	x := 25
	y := &x // Estoy pasando el valor de la direccion
	
	fmt.Println("Imprimiendo dir &x:", &x)
	fmt.Println("Imprimiendo Dir y: ", y)     // dir de memorioa
	fmt.Println("Imprimiendo Valor *y: ", *y) // *Y valor
	fmt.Println("Valor de x:", x)
	cambiarValor(x)
	fmt.Println("Valor de x:", x)
}

/**
* Cuando se pasa por valor, s
* cambiaValor(x) solo se está pasando el valor de x
*
* Por Referencia
* cambiaValor(&x) se envia el apuntador del la variable x
 */
func cambiarValor(a int) {
	a = 36
	fmt.Println("Valor a: ", a)
}
